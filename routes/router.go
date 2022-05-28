package routes

import (
	"Atlantis/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(middlewares.CORSMiddleware())

	//API Route Group
	v1 := router.Group("/api")
	v1Routes(v1)

	return router
}
