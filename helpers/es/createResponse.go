package helpers

import (
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs/requests"
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/olivere/elastic/v7"
)

func CreateResponse(ctx context.Context, ResponseData *requests.Response, sentryCtx context.Context) (string, error) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddResponse")
	defer span.Finish()

	if ResponseData.ResponseId != " " && ResponseData.Status {
		dbSpan1 := sentry.StartSpan(span.Context(), "[DB] update responses")
		multiMatchQuery, err := es.Client().Update().Index("responses").Id(ResponseData.ResponseId).Script(elastic.NewScriptInline(`ctx._source.Status = true`)).
			Upsert(map[string]interface{}{"retweets": 0}).
			Do(ctx)

		dbSpan1.Finish()

		if err != nil {
			fmt.Println(err)
			sentry.CaptureException(err)
			logger.Client().Error(err.Error())
			return "null", err
		}

		return multiMatchQuery.Id, nil
	} else {

		dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /responses")
		multiMatchQuery, err := es.Client().Index().Index("responses").BodyJson(ResponseData).Do(ctx)
		dbSpan1.Finish()

		if err != nil {
			fmt.Println(err)
			sentry.CaptureException(err)
			logger.Client().Error(err.Error())
			return "null", err
		}

		return multiMatchQuery.Id, nil
	}

}
