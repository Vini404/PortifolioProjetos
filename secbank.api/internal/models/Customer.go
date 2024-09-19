package models

import "time"

type Customer struct {
	ID               int
	FullName         string
	Phone            string
	Email            string
	Birthday         time.Time
	CreatedTimeStamp time.Time
	UpdatedTimeStamp time.Time
}
