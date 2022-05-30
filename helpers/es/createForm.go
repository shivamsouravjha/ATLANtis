package helpers

import (
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs/requests"
	"context"

	"github.com/getsentry/sentry-go"
)

func CreateForm(ctx context.Context, FormData *requests.Form, formID string, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddForm")
	defer span.Finish()

	if FormData.Name == "" {
		FormData.Name = "No Title"
	}

	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /forms")
	_, err := es.Client().Index().Id(formID).Index("forms").BodyJson(FormData).Do(ctx)
	dbSpan1.Finish()

	if err != nil {
		sentry.CaptureException(err)
		logger.Client().Error(err.Error())
		return
	}

}
