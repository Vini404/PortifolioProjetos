package models

import "time"

type Customer struct {
	ID               int       `db:"id"`
	FullName         string    `db:"fullname"`
	Phone            string    `db:"phone"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	Document         string    `db:"document"`
	Birthday         time.Time `db:"birthday"`
	CreatedTimeStamp time.Time `db:"createdtimestamp"`
	UpdatedTimeStamp time.Time `db:"updatedtimestamp"`
}
