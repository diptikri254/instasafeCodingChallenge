package router

import (
	"context"
	"instasafeCodingChallenge/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	transactionHandler transaction.Handler
}

func NewApi(transactionHandler transaction.Handler) *api {
	return &api{
		transactionHandler: transactionHandler,
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
		locationRouter.POST("/set")
		locationRouter.POST("/reset")
	}
}
