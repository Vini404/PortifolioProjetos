package dto

import "time"

type CreateCustomerRequest struct {
	FullName string    `json:"fullname"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string    `json:"password"`
}
