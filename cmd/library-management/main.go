package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitialPage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "This is a library management."})
}

func main() {
	router := gin.Default()

	//grouping routes at /api
	api := router.Group("/api")
	{
		api.GET("/initial", InitialPage)
	}

	router.Run()
}
