package dto

import "time"

type CreateAccountRequest struct {
	FullName string    `json:"fullname"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string    `json:"password"`
}
