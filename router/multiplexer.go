package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIMultiplexer(c context.Context) {

	routerVar.Use(gin.Logger(), gin.Recovery())

	routerVar.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	transactionRouter := routerVar.Group("/transactions")
	{
		transactionRouter.POST("")
		transactionRouter.DELETE("")
	}

	routerVar.GET("/statistics")

	locationRouter := routerVar.Group("/location")
	{
		locationRouter.POST("/set")
		locationRouter.POST("/reset")
	}
}
