package routers

import (
	"assignment_3/controllers"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	router.GET("/status/", controllers.Status)
	return router
}
