package POST

import (
	"Atlantis/constants"
	helpers "Atlantis/helpers/es"
	"Atlantis/services/logger"
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
	formID, err := helpers.CreateAnswer(ctx, &answerRequest, span.Context())
	if err != nil {
		resp.Status = constants.API_FAILED_STATUS
		resp.Message = err.Error()
		logger.Client().Error(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Creator updated successfully"
	resp.Data = formID
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
