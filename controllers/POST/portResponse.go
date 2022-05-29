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
	formID, err := helpers.CreateResponse(ctx, &responseRequest, span.Context())
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
