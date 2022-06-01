package routes

import (
	"Atlantis/controllers/GET"
	"Atlantis/controllers/POST"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func v1Routes(route *gin.RouterGroup) {

	router := gin.New()
	router.Use(apmgin.Middleware(router))
	// router.Use(gin.Logger())

	v1Routes := route.Group("/v1")
	{
		v1Routes.POST("/createForm", POST.CreateFormHandler)
		v1Routes.POST("/createQuestion", POST.CreateQuestionHandler)
		v1Routes.POST("/createResponse", POST.CreateResponseHandler)
		v1Routes.POST("/createAnswer", POST.CreateAnswerHandler)
		v1Routes.GET("/getForm", GET.GetFormHandler)
		v1Routes.GET("/createQuestion", GET.GetQuestionHandler)
		v1Routes.GET("/getResponse", GET.GetResponseHandler)
		v1Routes.GET("/createAnswer", GET.GetAnswerHandler)

	}
}
