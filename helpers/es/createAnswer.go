package helpers

import (
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs/requests"
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
)

func CreateAnswer(ctx context.Context, AnswerData *requests.Answer, sentryCtx context.Context) (string, error) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddAnswer")
	defer span.Finish()

	if AnswerData.AnswerID != "" {
		dbSpan1 := sentry.StartSpan(span.Context(), "[DB] update answer")
		multiMatchQuery, err := es.Client().Update().Index("answers").Id(AnswerData.AnswerID).Doc(map[string]interface{}{"Answer": AnswerData.Answer}).Do(ctx)

		dbSpan1.Finish()

		if err != nil {
			fmt.Println(err)
			sentry.CaptureException(err)
			logger.Client().Error(err.Error())
			return "null", err
		}

		return multiMatchQuery.Id, nil
	} else {

		dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /answer")
		multiMatchQuery, err := es.Client().Index().Index("answers").BodyJson(AnswerData).Do(ctx)
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
