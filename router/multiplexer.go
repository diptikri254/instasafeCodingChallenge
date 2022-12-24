package router

import (
	"context"
	"instasafeCodingChallenge/location"
	"instasafeCodingChallenge/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	transactionHandler transaction.Handler
	locationHandler    location.Handler
}

func NewApi(transactionHandler transaction.Handler,
	locationHandler location.Handler) *api {
	return &api{
		transactionHandler: transactionHandler,
		locationHandler:    locationHandler,
	}
}

func APIMultiplexer(c context.Context, apiSvc *api) {

	routerVar.Use(gin.Logger(), gin.Recovery())

	routerVar.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	transactionRouter := routerVar.Group("/transactions")
	{
		transactionRouter.POST("", apiSvc.transactionHandler.AddTransaction())
		transactionRouter.DELETE("", apiSvc.transactionHandler.DeleteTransactions())
	}

	routerVar.GET("/statistics")

	locationRouter := routerVar.Group("/location")
	{
		locationRouter.POST("/set", apiSvc.locationHandler.SetLocation())
		locationRouter.POST("/reset", apiSvc.locationHandler.ResetLocation())
	}
}
