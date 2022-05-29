package helpers

import (
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs/requests"
	"context"

	"github.com/getsentry/sentry-go"
)

func CreateAnswer(ctx context.Context, QuestionData *requests.Question, sentryCtx context.Context) (string, error) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddQuestion")
	defer span.Finish()

	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /questions")
	QuestionInsert, err := es.Client().Index().Index("questions").BodyJson(QuestionData).Do(ctx)
	dbSpan1.Finish()

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return "null", err
	}

	return QuestionInsert.Id, nil
}
