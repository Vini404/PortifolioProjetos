package models

import "time"

type AccountHolder struct {
	ID               int       `db:"id"`
	IDCustomer       int       `db:"idcustomer"`
	IsActive         bool      `db:"isactive"`
	CreatedTimeStamp time.Time `db:"createdtimestamp"`
	UpdatedTimeStamp time.Time `db:"updatedtimestamp"`
}
