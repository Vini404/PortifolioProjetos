package mock_interfaces

import (
	"github.com/stretchr/testify/mock"
	dto "secbank.api/dto/account"
	"secbank.api/models"
)

type MockIAccountService struct {
	mock.Mock
}

func (m *MockIAccountService) S_List() (*[]models.Account, error) {
	args := m.Called()
	if accounts, ok := args.Get(0).(*[]models.Account); ok {
		return accounts, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockIAccountService) S_Create(account models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockIAccountService) S_Update(account models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockIAccountService) S_Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockIAccountService) S_Get(id int) (*models.Account, error) {
	args := m.Called(id)
	if account, ok := args.Get(0).(*models.Account); ok {
		return account, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockIAccountService) S_GetInformationAccount(id int) (*dto.InformationAccountResponse, error) {
	args := m.Called(id)
	if info, ok := args.Get(0).(*dto.InformationAccountResponse); ok {
		return info, args.Error(1)
	}
	return nil, args.Error(1)
}
