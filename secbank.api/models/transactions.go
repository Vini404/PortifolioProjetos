package models

import "time"

type Transaction struct {
	ID               int       `db:"id"`
	IDDebitAccount   int       `db:"iddebitaccount"`
	IDCreditAccount  int       `db:"idcreditaccount"`
	Amount           float64   `db:"amount"`
	CreatedTimeStamp time.Time `db:"createdtimestamp"`
	TransactionType  int       `db:"transactiontype"`
}
