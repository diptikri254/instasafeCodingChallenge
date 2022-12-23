package transaction

import "time"

type DbService interface {
	AddTransaction(transaction Transaction)
	GetTransactions(d time.Duration) []Transaction
	DeleteTransactions()
}

type transactionDbService struct {
}

func (db *transactionDbService) AddTransaction(transaction Transaction) {
	transactionsList = append(transactionsList, transaction)
}

func (db *transactionDbService) GetTransactions(d time.Duration) []Transaction {
	var response []Transaction
	for _, v := range transactionsList {
		if v.TimeStamp.Add(d).After(time.Now()) {
			response = append(response, v)
		}
	}
	return response
}

func (db *transactionDbService) DeleteTransactions() {
	transactionsList = make([]Transaction, 0)
}

func NewTransactionDbService() DbService {
	return &transactionDbService{}
}
