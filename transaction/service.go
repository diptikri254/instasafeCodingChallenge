package transaction

import "time"

type Service interface {
	AddTransaction(transaction Transaction)
	GetTransactions(d time.Duration) []Transaction
	DeleteTransactions()
}

type transactionService struct {
	dbService DbService
}

func (ts *transactionService) AddTransaction(transaction Transaction) {
	ts.dbService.AddTransaction(transaction)
}

func (ts *transactionService) GetTransactions(d time.Duration) []Transaction {
	return ts.dbService.GetTransactions(d)
}

func (ts *transactionService) DeleteTransactions() {
	ts.dbService.DeleteTransactions()
}

func NewTransactionService(dbService DbService) Service {
	return &transactionService{
		dbService: dbService,
	}
}
