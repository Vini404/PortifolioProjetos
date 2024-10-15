package dto

import (
	"time"
)

type Response struct {
	Success      bool        `json:"success"`
	MessageError string      `json:"messageError"`
	Data         interface{} `json:"data"`
	StatusCode   int         `json:"statusCode"`
	Timestamp    time.Time   `json:"timestamp"`
}
