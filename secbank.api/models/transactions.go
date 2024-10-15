package models

import "time"

type Transaction struct {
	ID               int       `db:"id"`
	IDDebitAccount   int       `db:"id_debit_account"`
	IDCreditAccount  int       `db:"id_credit_account"`
	Amount           float64   `db:"amount"`
	CreatedTimeStamp time.Time `db:"created_time_stamp"`
	TransactionType  int       `db:"transaction_type"`
}
