package mock_interfaces

import (
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	dto "secbank.api/dto/transaction"
)

type MockITransactionService struct {
	mock.Mock
}

func (m *MockITransactionService) Transfer(transferRequest dto.TransferRequest, file multipart.File) error {
	args := m.Called(transferRequest, file)
	return args.Error(0)
}
