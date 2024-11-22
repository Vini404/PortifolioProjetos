package controllers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"secbank.api/controllers"
	dto "secbank.api/dto/balance"
	mock_interfaces "secbank.api/mocks"
	"secbank.api/models"
	"testing"
)

func TestBalanceController_Get(t *testing.T) {
	mockService := new(mock_interfaces.MockIBalanceService)
	controller := &controllers.BalanceController{IBalanceService: mockService}

	// Dados esperados
	expectedBalance := &models.Balance{
		IDAccount: 1,
		Amount:    1500.50,
	}
	accountID := 1

	// Configuração do mock
	mockService.On("S_GetByAccountID", accountID).Return(expectedBalance, nil)

	// Configuração da requisição
	req := httptest.NewRequest(http.MethodGet, "/balances/1", nil)
	req = addChiURLParam(req, "accountID", "1")
	rec := httptest.NewRecorder()

	// Execução do método
	controller.Get(rec, req)

	// Verificação do status HTTP
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse da resposta JSON
	var response struct {
		Success bool            `json:"success"`
		Data    *models.Balance `json:"data"`
	}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verificação do payload
	assert.Equal(t, expectedBalance.Amount, response.Data.Amount)
}

func TestBalanceController_Extract(t *testing.T) {
	mockService := new(mock_interfaces.MockIBalanceService)
	controller := &controllers.BalanceController{IBalanceService: mockService}

	// Dados esperados
	expectedExtract := []*dto.ExtractResponse{
		{OperationName: "Deposit", Amount: 500.00, TransferType: "credit"},
		{OperationName: "Withdraw", Amount: 100.00, TransferType: "debit"},
		{OperationName: "Transfer", Amount: 200.00, TransferType: "credit"},
	}
	accountID := 1

	// Configuração do mock
	mockService.On("S_Extract", accountID).Return(expectedExtract, nil)

	// Configuração da requisição
	req := httptest.NewRequest(http.MethodGet, "/balances/1/extract", nil)
	req = addChiURLParam(req, "accountID", "1")
	rec := httptest.NewRecorder()

	// Execução do método
	controller.Extract(rec, req)

	// Verificação do status HTTP
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse da resposta JSON
	var response struct {
		Success bool                   `json:"success"`
		Data    []*dto.ExtractResponse `json:"data"`
	}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verificação do payload
	assert.Equal(t, expectedExtract[0].OperationName, response.Data[0].OperationName)
	assert.Equal(t, expectedExtract[0].Amount, response.Data[0].Amount)
	assert.Equal(t, expectedExtract[0].TransferType, response.Data[0].TransferType)

	assert.Equal(t, expectedExtract[1].OperationName, response.Data[1].OperationName)
	assert.Equal(t, expectedExtract[1].Amount, response.Data[1].Amount)
	assert.Equal(t, expectedExtract[1].TransferType, response.Data[1].TransferType)

	assert.Equal(t, expectedExtract[2].OperationName, response.Data[2].OperationName)
	assert.Equal(t, expectedExtract[2].Amount, response.Data[2].Amount)
	assert.Equal(t, expectedExtract[2].TransferType, response.Data[2].TransferType)
}
