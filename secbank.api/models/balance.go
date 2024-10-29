package models

import "time"

type Balance struct {
	ID               int       `db:"id"`
	IDAccount        int       `db:"idaccount"`
	Amount           float64   `db:"amount"`
	AmountBlocked    float64   `db:"amountblocked"`
	UpdatedTimeStamp time.Time `db:"updatedtimestamp"`
}
