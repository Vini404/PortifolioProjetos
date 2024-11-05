package models

import "time"

type Account struct {
	ID               int       `db:"id"`
	IDAccountHolder  int       `db:"idaccountholder"`
	IsActive         bool      `db:"isactive"`
	Number           string    `db:"number"`
	Digit            string    `db:"digit"`
	Description      string    `db:"description"`
	CreatedTimeStamp time.Time `db:"createdtimestamp"`
	UpdatedTimeStamp time.Time `db:"updatedtimestamp"`
}
