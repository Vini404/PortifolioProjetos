package test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	dto "secbank.api/dto/account"
	mock_interfaces "secbank.api/mocks"
	"secbank.api/models"
	"secbank.api/services"
	"testing"
)

func TestAccountService_S_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountRepository(ctrl)
	service := &services.AccountService{IAccountRepository: mockRepo}

	// Mockando o retorno
	expectedAccounts := &[]models.Account{
		{ID: 1, Number: "123", Digit: "4", IsActive: true},
		{ID: 2, Number: "456", Digit: "7", IsActive: false},
	}
	mockRepo.EXPECT().R_List().Return(expectedAccounts, nil)

	// Executando o teste
	accounts, err := service.S_List()
	assert.NoError(t, err)
	assert.Equal(t, expectedAccounts, accounts)
}

func TestAccountService_S_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountRepository(ctrl)
	service := &services.AccountService{IAccountRepository: mockRepo}

	account := models.Account{Number: "123", Digit: "4", IsActive: true}

	// Cenário de sucesso
	mockRepo.EXPECT().R_Create(account).Return(1, nil)

	err := service.S_Create(account)
	assert.NoError(t, err)

	// Cenário de erro
	mockRepo.EXPECT().R_Create(account).Return(0, errors.New("database error"))

	err = service.S_Create(account)
	assert.Error(t, err)
}

func TestAccountService_S_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountRepository(ctrl)
	service := &services.AccountService{IAccountRepository: mockRepo}

	account := models.Account{ID: 1, Number: "123", Digit: "4", IsActive: true}

	// Cenário de sucesso
	mockRepo.EXPECT().R_Update(account).Return(nil)

	err := service.S_Update(account)
	assert.NoError(t, err)

	// Cenário de erro
	mockRepo.EXPECT().R_Update(account).Return(errors.New("database error"))

	err = service.S_Update(account)
	assert.Error(t, err)
}

func TestAccountService_S_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountRepository(ctrl)
	service := &services.AccountService{IAccountRepository: mockRepo}

	// Cenário de sucesso
	mockRepo.EXPECT().R_Delete(1).Return(nil)

	err := service.S_Delete(1)
	assert.NoError(t, err)

	// Cenário de erro
	mockRepo.EXPECT().R_Delete(1).Return(errors.New("database error"))

	err = service.S_Delete(1)
	assert.Error(t, err)
}

func TestAccountService_S_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountRepository(ctrl)
	service := &services.AccountService{IAccountRepository: mockRepo}

	expectedAccount := &models.Account{ID: 1, Number: "123", Digit: "4", IsActive: true}

	// Cenário de sucesso
	mockRepo.EXPECT().R_Get(1).Return(expectedAccount, nil)

	account, err := service.S_Get(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, account)

	// Cenário de erro
	mockRepo.EXPECT().R_Get(1).Return(nil, errors.New("not found"))

	account, err = service.S_Get(1)
	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestAccountService_S_GetInformationAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountRepository(ctrl)
	service := &services.AccountService{IAccountRepository: mockRepo}

	expectedInfo := &dto.InformationAccountResponse{
		AccountNumber: "123-4",
		CustomerName:  "John Doe",
		CustomerID:    1,
		IDAccount:     "1",
	}

	// Cenário de sucesso
	mockRepo.EXPECT().R_GetInformationAccount(1).Return(expectedInfo, nil)

	info, err := service.S_GetInformationAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedInfo, info)

	// Cenário de erro
	mockRepo.EXPECT().R_GetInformationAccount(1).Return(nil, errors.New("not found"))

	info, err = service.S_GetInformationAccount(1)
	assert.Error(t, err)
	assert.Nil(t, info)
}
