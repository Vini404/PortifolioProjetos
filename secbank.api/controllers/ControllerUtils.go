package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"secbank.api/dto"
	"time"
)

func SetResponseError(res http.ResponseWriter, error error) {
	response := dto.Response{}

	response = dto.Response{
		Success:      false,
		Timestamp:    time.Now(),
		StatusCode:   400,
		MessageError: error.Error(),
	}

	log.Println(error.Error())

	SetResponse(res, response)
}

func SetResponseSuccess(res http.ResponseWriter) {
	response := dto.Response{
		Success:    true,
		Timestamp:  time.Now(),
		StatusCode: 200,
	}

	SetResponse(res, response)
}

func SetResponseSuccessWithPayload(res http.ResponseWriter, data interface{}) {
	response := dto.Response{
		Success:    true,
		Timestamp:  time.Now(),
		StatusCode: 200,
		Data:       data,
	}

	SetResponse(res, response)
}

func SetResponse(res http.ResponseWriter, response dto.Response) {
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(response.StatusCode)

	json.NewEncoder(res).Encode(response)
}
