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
	ctx := c.Request.Context()
	resp := response.EventResponse{}
	formID, err := helpers.CreateForm(ctx, &formRequest, span.Context())
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
