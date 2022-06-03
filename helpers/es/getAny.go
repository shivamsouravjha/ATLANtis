package helpers

import (
	"Atlantis/constants"
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs/requests"

	"context"

	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

func GetAny(ctx context.Context, AnyData *requests.AnyHandler, sentryCtx context.Context) ([]interface{}, error) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddForm")
	defer span.Finish()

	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /forms")
	res, err := es.Client().Search().Index(constants.IndexElasticSearch[AnyData.Index]).Query(QueryDetails(AnyData.Param, AnyData.Value)).Size(1000).
		FetchSourceContext(elastic.NewFetchSourceContext(true)).Do(ctx)
	dbSpan1.Finish()

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return []interface{}{}, err
	}
	var data1 interface{}
	var dataRes []interface{}
	if res != nil {
		for _, s := range res.Hits.Hits {
			jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(s.Source, &data1)
			dataRes = append(dataRes, data1)
		}
	}
	return dataRes, nil
}

func QueryDetails(param string, value string) *elastic.TermQuery {
	return elastic.NewTermQuery(param, value)
}
