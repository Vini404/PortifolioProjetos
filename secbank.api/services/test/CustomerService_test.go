package test

import (
	"bytes"
	"mime/multipart"
	dto "secbank.api/dto/account"
	dto3 "secbank.api/dto/balance"
	dto2 "secbank.api/dto/customer"
	"secbank.api/models"
	"secbank.api/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para CustomerRepository
type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) R_List() (*[]models.Customer, error) {
	args := m.Called()
	return args.Get(0).(*[]models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) R_Create(customer models.Customer) (int, error) {
	args := m.Called(customer)
	return args.Int(0), args.Error(1)
}

func (m *MockCustomerRepository) R_Get(id int) (*models.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) R_Get_By_Email(email string) (*models.Customer, error) {
	args := m.Called(email)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockCustomerRepository) R_Update(customer models.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) R_Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Mock para AccountHolderRepository
type MockAccountHolderRepository struct {
	mock.Mock
}

func (m *MockAccountHolderRepository) R_Get(id int) (*models.AccountHolder, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockAccountHolderRepository) R_Create(holder models.AccountHolder) (int, error) {
	args := m.Called(holder)
	return args.Int(0), args.Error(1)
}

func (m *MockAccountHolderRepository) R_List() (*[]models.AccountHolder, error) {
	args := m.Called()
	return args.Get(0).(*[]models.AccountHolder), args.Error(1)
}

func (m *MockAccountHolderRepository) R_Update(accountHolder models.AccountHolder) error {
	args := m.Called(accountHolder)
	return args.Error(0)
}

func (m *MockAccountHolderRepository) R_Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Mock para AccountRepository
type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) R_Get(id int) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockAccountRepository) R_Get_By_Number_And_Digit(number string, digit string) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockAccountRepository) R_GetInformationAccount(id int) (*dto.InformationAccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockAccountRepository) R_GetAccountByCustomer(customerID int) (*models.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockAccountRepository) R_Create(account models.Account) (int, error) {
	args := m.Called(account)
	return args.Int(0), args.Error(1)
}

func (m *MockAccountRepository) R_List() (*[]models.Account, error) {
	args := m.Called()
	return args.Get(0).(*[]models.Account), args.Error(1)
}

func (m *MockAccountRepository) R_Update(account models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) R_Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Mock para BalanceRepository
type MockBalanceRepository struct {
	mock.Mock
}

func (m *MockBalanceRepository) R_Create(balance models.Balance) (int, error) {
	args := m.Called(balance)
	return args.Int(0), args.Error(1)
}

func (m *MockBalanceRepository) R_GetByAccountID(accountID int) (*models.Balance, error) {
	args := m.Called(accountID)
	return args.Get(0).(*models.Balance), args.Error(1)
}

func (m *MockBalanceRepository) R_Extract(accountID int) ([]*dto3.ExtractResponse, error) {
	args := m.Called(accountID)
	return args.Get(0).([]*dto3.ExtractResponse), args.Error(1)
}

func (m *MockBalanceRepository) R_Update(balance *models.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

// Função auxiliar para criar um arquivo simulado
type mockMultipartFile struct {
	*bytes.Reader
}

func (m *mockMultipartFile) Close() error {
	return nil
}

func createMockFile(content string) multipart.File {
	return &mockMultipartFile{Reader: bytes.NewReader([]byte(content))}
}

// Teste para S_Create
func TestCustomerService_S_Create(t *testing.T) {
	mockCustomerRepo := &MockCustomerRepository{}
	mockAccountHolderRepo := &MockAccountHolderRepository{}
	mockAccountRepo := &MockAccountRepository{}
	mockBalanceRepo := &MockBalanceRepository{}

	// Configuração dos mocks
	mockCustomerRepo.On("R_Get_By_Email", "test@example.com").Return((*models.Customer)(nil), nil) // No existing user
	mockCustomerRepo.On("R_Create", mock.Anything).Return(1, nil)
	mockAccountHolderRepo.On("R_Create", mock.Anything).Return(1, nil)
	mockAccountRepo.On("R_Create", mock.Anything).Return(1, nil)
	mockBalanceRepo.On("R_Create", mock.Anything).Return(1, nil)

	service := services.CustomerService{
		ICustomerRepository:      mockCustomerRepo,
		IAccountHolderRepository: mockAccountHolderRepo,
		IAccountRepository:       mockAccountRepo,
		IBalanceRepository:       mockBalanceRepo,
	}

	file := createMockFile("fake image data")
	customer := models.Customer{
		Email:    "test@example.com",
		Password: "password123",
		Phone:    "123456789",
	}

	service.S_Create(customer, file)
}

// Teste para S_Auth com sucesso
func TestCustomerService_S_Auth_Success(t *testing.T) {
	mockCustomerRepo := &MockCustomerRepository{}

	mockCustomerRepo.On("R_Get_By_Email", "test@example.com").Return(&models.Customer{
		ID:       1,
		Email:    "test@example.com",
		Password: "password123",
	}, nil)

	service := services.CustomerService{
		ICustomerRepository: mockCustomerRepo,
	}

	request := dto2.AuthRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	response, err := service.S_Auth(request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Token)

	mockCustomerRepo.AssertExpectations(t)
}

// Teste para S_Auth com senha inválida
func TestCustomerService_S_Auth_InvalidPassword(t *testing.T) {
	mockCustomerRepo := &MockCustomerRepository{}

	mockCustomerRepo.On("R_Get_By_Email", "test@example.com").Return(&models.Customer{
		ID:       1,
		Email:    "test@example.com",
		Password: "wrongpassword",
	}, nil)

	service := services.CustomerService{
		ICustomerRepository: mockCustomerRepo,
	}

	request := dto2.AuthRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	response, err := service.S_Auth(request)
	assert.Error(t, err)
	assert.Nil(t, response)

	mockCustomerRepo.AssertExpectations(t)
}

// Teste para S_Get
func TestCustomerService_S_Get(t *testing.T) {
	mockCustomerRepo := &MockCustomerRepository{}

	expectedCustomer := &models.Customer{
		ID:    1,
		Email: "test@example.com",
	}

	mockCustomerRepo.On("R_Get", 1).Return(expectedCustomer, nil)

	service := services.CustomerService{
		ICustomerRepository: mockCustomerRepo,
	}

	customer, err := service.S_Get(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	mockCustomerRepo.AssertExpectations(t)
}
