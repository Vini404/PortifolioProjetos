package models

import "time"

type Balance struct {
	ID               int       `db:"id"`
	IDAccount        int       `db:"id_account"`
	Amount           float64   `db:"amount"`
	AmountBlocked    float64   `db:"amount_blocked"`
	UpdatedTimeStamp time.Time `db:"updated_time_stamp,omitempty"`
}
