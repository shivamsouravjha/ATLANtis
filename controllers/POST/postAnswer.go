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

func CreateAnswerHandler(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] AddAnswerHandler", sentry.TransactionName("Create Answer Handler"))
	defer span.Finish()

	answerRequest := requests.Answer{}
	if err := c.ShouldBind(&answerRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	ctx := c.Request.Context()
	resp := response.EventResponse{}
	var answerId string
	if answerRequest.AnswerID != "" {
		answerId = answerRequest.FormID
	} else {
		answerId = utils.GeneratorUUID(11)
	}

	helpers.CreateAnswer(ctx, &answerRequest, answerId, span.Context())

	resp.Status = "Success"
	resp.Message = "Creator updated successfully"
	resp.Data = answerId
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
