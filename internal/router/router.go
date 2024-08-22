package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func to config routes
func SetRouter() *gin.Engine {
	router := gin.Default()

	//grouping routes at /api
	api := router.Group("/api")
	{
		api.GET("/initial", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "This is a library management."})
		})
	}

	//NoRoute -> adds handlers for NoRoute
	router.NoRoute()

	return router
}
