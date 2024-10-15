package models

import "time"

type AccountHolder struct {
	ID               int       `db:"id"`
	IDCustomer       int       `db:"id_customer"`
	IsActive         bool      `db:"is_active"`
	CreatedTimeStamp time.Time `db:"created_time_stamp"`
	UpdatedTimeStamp time.Time `db:"updated_time_stamp,omitempty"`
}
