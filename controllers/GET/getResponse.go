package GET

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

func GetResponseHandler(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] AddResponseHandler", sentry.TransactionName("Create Response Handler"))
	defer span.Finish()

	responseRequest := requests.GetResponse{}
	if err := c.ShouldBind(&responseRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}
	ctx := c.Request.Context()
	resp := response.RespResponse{}
	isUpdate := true

	responses, err := helpers.GetResponse(ctx, &responseRequest, isUpdate, span.Context())
	if err != nil {
		resp.Status = "Failed"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Status = "Success"
	resp.Message = "Creator updated successfully"
	resp.Data = responses
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
