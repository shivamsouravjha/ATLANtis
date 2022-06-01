package helpers

import (
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs/requests"
	"Atlantis/structs/response"
	"context"
	"encoding/json"
	"fmt"

	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

func GetResponse(ctx context.Context, ResponseData *requests.GetResponse, isUpdate bool, sentryCtx context.Context) (response.Response, error) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddResponse")
	defer span.Finish()

	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /forms")
	res, err := es.Client().Search().Index("responses").Query(ResponseDetails(ResponseData.ResponseId)).Size(1).
		FetchSourceContext(elastic.NewFetchSourceContext(true)).Do(ctx)
	dbSpan1.Finish()

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return response.Response{}, err
	}
	var data1 requests.Response
	if res != nil {
		for _, s := range res.Hits.Hits {
			err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(s.Source, &data1)
			fmt.Println(data1)
		}
	}

	dbSpan2 := sentry.StartSpan(span.Context(), "[DB] Insert into /forms")
	res2, err := es.Client().Search().Index("answers").SearchSource(elastic.NewSearchSource().Query(elastic.NewMatchQuery("responseId", ResponseData.ResponseId)).Size(1000)).Size(1000).Do(ctx)
	rescfg, _ := json.Marshal(elastic.NewSearchSource().Query(elastic.NewMatchQuery("answerId", "sJeoD4EBP9dta9N7JaUi")).Size(1000))
	fmt.Println(string(rescfg))
	dbSpan2.Finish()

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return response.Response{}, err
	}
	var data2 []requests.Answer
	var temp requests.Answer

	if res2 != nil {
		for _, s := range res2.Hits.Hits {
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(s.Source, &temp)
			data2 = append(data2, temp)
		}
	}

	data := response.Response{
		Response: data1,
		Answer:   data2,
	}
	return data, nil

}
func ResponseDetails(id string) *elastic.TermQuery {
	return elastic.NewTermQuery("_id", id)
}
