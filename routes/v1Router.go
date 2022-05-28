package routes

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func v1Routes(route *gin.RouterGroup) {

	router := gin.New()
	router.Use(apmgin.Middleware(router))
	// router.Use(gin.Logger())

	v1Routes := route.Group("/v1")
	{
		v1Routes.GET("/healthCheck")
	}
}
