package models

import "time"

type Account struct {
	ID               int       `db:"id"`
	IDAccountHolder  int       `db:"id_account_holder"`
	IsActive         bool      `db:"is_active"`
	Number           string    `db:"number"`
	Digit            string    `db:"digit"`
	CreatedTimeStamp time.Time `db:"created_time_stamp"`
	UpdatedTimeStamp time.Time `db:"updated_time_stamp,omitempty"`
}
