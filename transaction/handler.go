package transaction

import (
	"log"
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
		type Request struct {
			Amount    float64 `json:"amount"`
			TimeStamp string  `json:"timestamp"`
		}
		var requestBody Request
		err := c.ShouldBind(&requestBody)
		if err != nil {
			log.Printf("AddTransaction: error while decoding requestbody: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid json",
			})
			return
		}

		timeStamp, err := time.Parse(time.RFC3339, requestBody.TimeStamp)
		if err != nil {
			log.Printf("AddTransaction: invalid timestamp: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid timestamp",
			})
			return
		}

		transaction := Transaction{
			Amount:    requestBody.Amount,
			TimeStamp: timeStamp,
		}

		if transaction.TimeStamp.After(time.Now()) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "unprocessable entity",
			})
			return
		}

		handler.transactionService.AddTransaction(transaction)

		if !transaction.TimeStamp.After(time.Now().Add(-1 * time.Minute)) {
			c.JSON(http.StatusNoContent, gin.H{})
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
