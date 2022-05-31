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

func CreateResponseHandler(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] AddResponseHandler", sentry.TransactionName("Create Response Handler"))
	defer span.Finish()

	responseRequest := requests.Response{}
	if err := c.ShouldBind(&responseRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	ctx := c.Request.Context()
	resp := response.EventResponse{}

	if responseRequest.ResponseId != "" {
		responseRequest.ResponseId = utils.GeneratorUUID(11)
	}
	go helpers.CreateResponse(ctx, &responseRequest, span.Context())

	resp.Status = "Success"
	resp.Message = "Creator updated successfully"
	resp.Data = responseRequest.ResponseId
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
