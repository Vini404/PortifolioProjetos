package controllers

import (
	"encoding/json"
	"net/http"
	"secbank.api/auth"
	dto "secbank.api/dto/transaction"
	interfaces "secbank.api/interfaces/service"
	"strings"
)

type TransactionController struct {
	interfaces.ITransactionService
}

func (controller *TransactionController) Transfer(res http.ResponseWriter, req *http.Request) {
	var transferUserRequest dto.TransferUserRequest

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&transferUserRequest)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	authHeader := req.Header.Get("Authorization")

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

	transferRequest := dto.TransferRequest{
		Amount:                  transferUserRequest.Amount,
		IDCreditAccount:         transferUserRequest.IDCreditAccount,
		IDCustomerOriginAccount: auth.GetCustomerIDByJwtToken(tokenString),
	}

	errTransfer := controller.ITransactionService.Transfer(transferRequest)

	if errTransfer != nil {
		SetResponseError(res, errTransfer)
		return
	}

	SetResponseSuccess(res)
}
