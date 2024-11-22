package mock_interfaces

import (
	"github.com/stretchr/testify/mock"
	dto "secbank.api/dto/balance"
	"secbank.api/models"
)

type MockIBalanceService struct {
	mock.Mock
}

func (m *MockIBalanceService) S_GetByAccountID(accountID int) (*models.Balance, error) {
	args := m.Called(accountID)
	if balance, ok := args.Get(0).(*models.Balance); ok {
		return balance, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockIBalanceService) S_Extract(accountID int) ([]*dto.ExtractResponse, error) {
	args := m.Called(accountID)
	if extract, ok := args.Get(0).([]*dto.ExtractResponse); ok {
		return extract, args.Error(1)
	}
	return nil, args.Error(1)
}
