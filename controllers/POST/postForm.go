package POST

import (
	helpers "Atlantis/helpers/es"
	"Atlantis/structs/requests"
	"Atlantis/structs/response"
	"Atlantis/utils"
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func CreateFormHandler(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] AddFormHandler", sentry.TransactionName("Create Form Handler"))
	defer span.Finish()

	formRequest := requests.Form{}
	if err := c.ShouldBind(&formRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}

	formID := utils.GeneratorUUID(11)
	ctx := c.Request.Context()
	resp := response.EventResponse{}
	helpers.CreateForm(ctx, &formRequest, formID, span.Context())

	resp.Status = "Success"
	resp.Message = "Creator updated successfully"
	resp.Data = formID
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
