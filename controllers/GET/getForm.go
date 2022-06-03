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

func GetFormHandler(c *gin.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(context.TODO(), "[GIN] GetFormHandler", sentry.TransactionName("Get Form Handler"))
	defer span.Finish()

	formRequest := requests.GetForm{}
	if err := c.ShouldBind(&formRequest); err != nil {
		span.Status = sentry.SpanStatusFailedPrecondition
		sentry.CaptureException(err)
		c.JSON(422, utils.SendErrorResponse(err))
		return
	}

	ctx := c.Request.Context()
	resp := response.FormResponse{}

	forms, err := helpers.GetForm(ctx, &formRequest, span.Context())
	if err != nil {
		resp.Status = "Failed"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Status = "Success"
	resp.Message = "Creator updated successfully"
	resp.Data = forms
	span.Status = sentry.SpanStatusOK

	c.JSON(http.StatusOK, resp)

}
