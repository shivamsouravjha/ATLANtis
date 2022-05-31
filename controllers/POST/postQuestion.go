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

func CreateQuestionHandler(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] CreateQuestionHandler", sentry.TransactionName("Create Question Handler"))
	defer span.Finish()

	questionRequest := requests.Question{}
	if err := c.ShouldBind(&questionRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	ctx := c.Request.Context()
	resp := response.EventResponse{}

	questionID := utils.GeneratorUUID(11)

	go helpers.CreateQuestion(ctx, &questionRequest, questionID, span.Context())

	resp.Status = "Success"
	resp.Message = "Creator updated successfully"
	resp.Data = questionID
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
