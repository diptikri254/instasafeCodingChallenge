package transaction

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddTransaction() gin.HandlerFunc
	DeleteTransactions() gin.HandlerFunc
}

type transactionHandler struct {
	transactionService Service
}

func (handler *transactionHandler) AddTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody Transaction
		err := c.ShouldBind(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid json",
			})
			return
		}

		if requestBody.TimeStamp.After(time.Now()) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "unprocessable entity",
			})
			return
		}

		handler.transactionService.AddTransaction(requestBody)

		if requestBody.TimeStamp.After(time.Now().Add(-1 * time.Minute)) {
			c.JSON(http.StatusNoContent, gin.H{
				"message": "no content",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "created",
		})
	}
}

func (handler *transactionHandler) DeleteTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.transactionService.DeleteTransactions()
		c.JSON(http.StatusNoContent, gin.H{
			"message": "deleted",
		})
	}
}

func NewTransactionHandler(transactionService Service) Handler {
	return &transactionHandler{
		transactionService: transactionService,
	}
}
