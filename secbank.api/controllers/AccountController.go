package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/jinzhu/copier"
	"net/http"
	"secbank.api/dto/account"
	"secbank.api/interfaces/service"
	"secbank.api/models"
	"strconv"
)

type AccountController struct {
	interfaces.IAccountService
}

func (controller *AccountController) List(res http.ResponseWriter, req *http.Request) {

	account, err := controller.S_List()

	if err != nil {
		SetResponseError(res, err)
		return
	}

	SetResponseSuccessWithPayload(res, account)
}

func (controller *AccountController) Create(res http.ResponseWriter, req *http.Request) {
	var accountRequest dto.CreateAccountRequest
	var account models.Account

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&accountRequest)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	copier.Copy(&account, &accountRequest)

	errInsert := controller.S_Create(account)

	if errInsert != nil {
		SetResponseError(res, errInsert)
		return
	}

	SetResponseSuccess(res)
}

func (controller *AccountController) Update(res http.ResponseWriter, req *http.Request) {
	var account models.Account
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&account)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	updateErr := controller.S_Update(account)

	if updateErr != nil {
		SetResponseError(res, updateErr)
		return
	}

	SetResponseSuccess(res)
}

func (controller *AccountController) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		SetResponseError(res, err)
		return
	}
	account, errGet := controller.S_Get(idParsed)

	if errGet != nil {
		SetResponseError(res, errGet)
		return
	}

	SetResponseSuccessWithPayload(res, account)
}

func (controller *AccountController) Delete(res http.ResponseWriter, req *http.Request) {
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
