package helpers

import (
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs/requests"
	"Atlantis/structs/response"
	"encoding/json"

	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

func GetForm(ctx context.Context, FormData *requests.GetForm, sentryCtx context.Context) (response.Form, error) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] GetForm")
	defer span.Finish()

	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Get from /forms")
	res, err := es.Client().Search().Index("forms").Query(FormDetails(FormData.FormID)).Size(1).
		FetchSourceContext(elastic.NewFetchSourceContext(true)).Do(ctx)
	dbSpan1.Finish()

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return response.Form{}, err
	}
	var data1 requests.Form
	if res != nil {
		for _, s := range res.Hits.Hits {
			err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(s.Source, &data1)
			fmt.Println(data1)
		}
	}

	dbSpan2 := sentry.StartSpan(span.Context(), "[DB] Get from /questions")
	res2, err := es.Client().Search().Index("questions").SearchSource(elastic.NewSearchSource().Query(elastic.NewMatchQuery("form", FormData.FormID)).Size(1000)).Size(1000).Do(ctx)
	rescfg, _ := json.Marshal(elastic.NewSearchSource().Query(elastic.NewMatchQuery("Form", "sJeoD4EBP9dta9N7JaUi")).Size(1000))
	fmt.Println(string(rescfg))
	dbSpan2.Finish()

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return response.Form{}, err
	}
	var data2 []requests.Question
	var temp requests.Question

	if res2 != nil {
		for _, s := range res2.Hits.Hits {
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(s.Source, &temp)
			data2 = append(data2, temp)
		}
	}

	data := response.Form{
		Form:     data1,
		Question: data2,
	}
	return data, nil
}

func FormDetails(id string) *elastic.TermQuery {
	return elastic.NewTermQuery("_id", id)
}
