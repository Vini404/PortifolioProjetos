package test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	dto "secbank.api/dto/balance"
	mock_interfaces "secbank.api/mocks"
	"secbank.api/models"
	"secbank.api/services"
)

func TestBalanceService_S_GetByAccountID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIBalanceRepository(ctrl)
	service := &services.BalanceService{IBalanceRepository: mockRepo}

	expectedBalance := &models.Balance{
		ID:               1,
		IDAccount:        1,
		Amount:           1000.50,
		AmountBlocked:    100.00,
		UpdatedTimeStamp: time.Now(),
	}

	// Cenário de sucesso
	mockRepo.EXPECT().R_GetByAccountID(1).Return(expectedBalance, nil)

	balance, err := service.S_GetByAccountID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, balance)

	// Cenário de erro
	mockRepo.EXPECT().R_GetByAccountID(1).Return(nil, errors.New("not found"))

	balance, err = service.S_GetByAccountID(1)
	assert.Error(t, err)
	assert.Nil(t, balance)
}

func TestBalanceService_S_Extract(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIBalanceRepository(ctrl)
	service := &services.BalanceService{IBalanceRepository: mockRepo}

	expectedExtract := []*dto.ExtractResponse{
		{OperationName: "Transferência", Amount: -200.00, TransferType: "Debito"},
		{OperationName: "Depósito", Amount: 500.00, TransferType: "Credito"},
	}

	// Cenário de sucesso
	mockRepo.EXPECT().R_Extract(1).Return(expectedExtract, nil)

	extract, err := service.S_Extract(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedExtract, extract)

	// Cenário de erro
	mockRepo.EXPECT().R_Extract(1).Return(nil, errors.New("not found"))

	extract, err = service.S_Extract(1)
	assert.Error(t, err)
	assert.Nil(t, extract)
}
