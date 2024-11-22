package routes

import (
	"github.com/gin-gonic/gin"
	"9ja-state/handlers"
)

func RegisterV1Routes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/states", handlers.GetAllStates)
		v1.GET("/states/:state", handlers.GetStateLGAs)
		v1.GET("/search", handlers.SearchStatesAndLGAs)
	}
}
