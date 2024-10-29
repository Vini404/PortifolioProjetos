package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	interfaces "secbank.api/interfaces/service"
	"strconv"
)

type BalanceController struct {
	interfaces.IBalanceService
}

func (controller *BalanceController) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "accountID")

	accountIDParsed, err := strconv.Atoi(id)

	if err != nil {
		SetResponseError(res, err)
		return
	}
	currentBalance, errGet := controller.S_GetByAccountID(accountIDParsed)

	if errGet != nil {
		SetResponseError(res, errGet)
		return
	}

	SetResponseSuccessWithPayload(res, currentBalance)
}

func (controller *BalanceController) Extract(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "accountID")

	accountIDParsed, err := strconv.Atoi(id)

	if err != nil {
		SetResponseError(res, err)
		return
	}
	extractResponse, errGet := controller.S_Extract(accountIDParsed)

	if errGet != nil {
		SetResponseError(res, errGet)
		return
	}

	SetResponseSuccessWithPayload(res, extractResponse)
}
