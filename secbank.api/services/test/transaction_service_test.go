package test

import (
	"secbank.api/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/stretchr/testify/mock"
)

// Mock para AccountRepository
type MockAccountRepository2 struct {
	mock.Mock
}

func (m *MockAccountRepository2) R_Get_By_Number_And_Digit(number, digit string) (*models.Account, error) {
	args := m.Called(number, digit)
	account, ok := args.Get(0).(*models.Account)
	if !ok && args.Get(0) != nil {
		panic("R_Get_By_Number_And_Digit: retorno inválido")
	}
	return account, args.Error(1)
}

func (m *MockAccountRepository2) R_GetAccountByCustomer(customerID int) (*models.Account, error) {
	args := m.Called(customerID)
	account, ok := args.Get(0).(*models.Account)
	if !ok && args.Get(0) != nil {
		panic("R_GetAccountByCustomer: retorno inválido")
	}
	return account, args.Error(1)
}

// Mock para BalanceRepository
type MockBalanceRepository2 struct {
	mock.Mock
}

func (m *MockBalanceRepository2) R_GetByAccountID(accountID int) (*models.Balance, error) {
	args := m.Called(accountID)
	balance, ok := args.Get(0).(*models.Balance)
	if !ok && args.Get(0) != nil {
		panic("R_GetByAccountID: retorno inválido")
	}
	return balance, args.Error(1)
}

func (m *MockBalanceRepository2) R_Update(balance *models.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

// Mock para TransactionRepository
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) R_Create(transaction models.Transaction) (int, error) {
	args := m.Called(transaction)
	return args.Int(0), args.Error(1)
}

// Mock para CustomerRepository
type MockCustomerRepository2 struct {
	mock.Mock
}

func (m *MockCustomerRepository2) R_Get(customerID int) (*models.Customer, error) {
	args := m.Called(customerID)
	customer, ok := args.Get(0).(*models.Customer)
	if !ok && args.Get(0) != nil {
		panic("R_Get: retorno inválido")
	}
	return customer, args.Error(1)
}

// Mock para RekognitionService
type MockRekognitionService struct {
	mock.Mock
}

func (m *MockRekognitionService) SearchUsersByImage(collectionID string, imageBytes []byte) (*rekognition.SearchUsersByImageOutput, error) {
	args := m.Called(collectionID, imageBytes)
	output, ok := args.Get(0).(*rekognition.SearchUsersByImageOutput)
	if !ok && args.Get(0) != nil {
		panic("SearchUsersByImage: retorno inválido")
	}
	return output, args.Error(1)
}

// Teste para a função Transfer
func TestTransactionService_Transfer(t *testing.T) {
	// Criação dos mocks
	mockAccountRepo := &MockAccountRepository{}
	mockBalanceRepo := &MockBalanceRepository{}
	mockTransactionRepo := &MockTransactionRepository{}
	mockCustomerRepo := &MockCustomerRepository{}

	// Configuração dos mocks
	mockAccountRepo.On("R_Get_By_Number_And_Digit", "123456", "0").Return(&models.Account{ID: 1}, nil)
	mockAccountRepo.On("R_GetAccountByCustomer", 2).Return(&models.Account{ID: 2}, nil)

	mockBalanceRepo.On("R_GetByAccountID", 2).Return(&models.Balance{Amount: 1000, AmountBlocked: 0}, nil)
	mockBalanceRepo.On("R_GetByAccountID", 1).Return(&models.Balance{Amount: 500, AmountBlocked: 0}, nil)

	mockBalanceRepo.On("R_Update", mock.Anything).Return(nil)

	mockTransactionRepo.On("R_Create", mock.Anything).Return(1, nil)

	mockCustomerRepo.On("R_Get", 2).Return(&models.Customer{ID: 2}, nil)

}
