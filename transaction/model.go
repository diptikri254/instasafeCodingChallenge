package transaction

import "time"

type Transaction struct {
	Amount    float64   `json:"amount"`
	TimeStamp time.Time `json:"timestamp"`
}

var transactionsList []Transaction
