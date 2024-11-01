package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/jinzhu/copier"
	"net/http"
	"secbank.api/auth"
	"secbank.api/dto/customer"
	"secbank.api/interfaces/service"
	"secbank.api/models"
	"strconv"
	"strings"
)

type CustomerController struct {
	interfaces.ICustomerService
}

func (controller *CustomerController) Auth(res http.ResponseWriter, req *http.Request) {
	var authRequest dto.AuthRequest

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&authRequest)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	token, err := controller.S_Auth(authRequest)

	if err != nil {
		SetResponseError(res, err)
		return
	}

	SetResponseSuccessWithPayload(res, token)

}

func (controller *CustomerController) List(res http.ResponseWriter, req *http.Request) {

	customers, err := controller.S_List()

	if err != nil {
		SetResponseError(res, err)
		return
	}

	SetResponseSuccessWithPayload(res, customers)
}

func (controller *CustomerController) Create(res http.ResponseWriter, req *http.Request) {
	var customerRequest dto.CreateCustomerRequest
	var customer models.Customer

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&customerRequest)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	copier.Copy(&customer, &customerRequest)

	errInsert := controller.S_Create(customer)

	if errInsert != nil {
		SetResponseError(res, errInsert)
		return
	}

	SetResponseSuccess(res)
}

func (controller *CustomerController) Update(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&customer)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	updateErr := controller.S_Update(customer)

	if updateErr != nil {
		SetResponseError(res, updateErr)
		return
	}

	SetResponseSuccess(res)
}

func (controller *CustomerController) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		SetResponseError(res, err)
		return
	}
	customer, errGet := controller.S_Get(idParsed)

	if errGet != nil {
		SetResponseError(res, errGet)
		return
	}

	SetResponseSuccessWithPayload(res, customer)
}

func (controller *CustomerController) GetCustomerByToken(res http.ResponseWriter, req *http.Request) {

	authHeader := req.Header.Get("Authorization")

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

	customer, errGet := controller.S_Get(auth.GetCustomerIDByJwtToken(tokenString))

	if errGet != nil {
		SetResponseError(res, errGet)
		return
	}

	SetResponseSuccessWithPayload(res, customer)
}

func (controller *CustomerController) Delete(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		SetResponseError(res, err)
		return
	}
	errToDelete := controller.S_Delete(idParsed)

	if errToDelete != nil {
		SetResponseError(res, errToDelete)
		return
	}

	SetResponseSuccess(res)
}
