package transaction

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	AddTransactions() gin.HandlerFunc
	DeleteTransactions() gin.HandlerFunc
}

type transactionHandler struct {
	transactionService *Service
}

func (handler *transactionHandler) AddTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func (handler *transactionHandler) DeleteTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func NewTransactionHandler(transactionService *Service) Handler {
	return &transactionHandler{
		transactionService: transactionService,
	}
}
