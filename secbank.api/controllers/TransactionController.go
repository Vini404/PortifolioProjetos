package controllers

import (
	"net/http"
	"secbank.api/auth"
	dto "secbank.api/dto/transaction"
	interfaces "secbank.api/interfaces/service"
	"strconv"
	"strings"
)

type TransactionController struct {
	interfaces.ITransactionService
}

func (controller *TransactionController) Transfer(res http.ResponseWriter, req *http.Request) {
	// Define um limite de memória para o arquivo de upload
	err := req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		SetResponseError(res, err)
		return
	}

	// Extrai o arquivo do formulário
	file, _, err := req.FormFile("file")
	if err != nil {
		SetResponseError(res, err)
		return
	}
	defer file.Close()

	// Extrai os valores dos outros campos do formulário
	amountStr := req.FormValue("Amount")
	digitCreditAccount := req.FormValue("DigitCreditAccount")   // Agora é uma string
	numberCreditAccount := req.FormValue("NumberCreditAccount") // Agora é uma string

	// Converte o campo Amount para float64
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		SetResponseError(res, err)
		return
	}

	// Extrai o token de autorização do cabeçalho
	authHeader := req.Header.Get("Authorization")
	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

	// Cria a estrutura de transferência com os dados extraídos
	transferRequest := dto.TransferRequest{
		Amount:                  amount,
		DigitCreditAccount:      digitCreditAccount,
		NumberCreditAccount:     numberCreditAccount,
		IDCustomerOriginAccount: auth.GetCustomerIDByJwtToken(tokenString),
	}

	// Chama o serviço de transferência
	errTransfer := controller.ITransactionService.Transfer(transferRequest, file)
	if errTransfer != nil {
		SetResponseError(res, errTransfer)
		return
	}

	SetResponseSuccess(res)
}
