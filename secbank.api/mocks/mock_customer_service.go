package mock_interfaces

import (
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	dto "secbank.api/dto/customer"
	"secbank.api/models"
)

type MockICustomerService struct {
	mock.Mock
}

func (m *MockICustomerService) S_List() (*[]models.Customer, error) {
	args := m.Called()
	if customers, ok := args.Get(0).(*[]models.Customer); ok {
		return customers, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockICustomerService) S_Create(customer models.Customer, file multipart.File) error {
	args := m.Called(customer, file)
	return args.Error(0)
}

func (m *MockICustomerService) S_Update(customer models.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockICustomerService) S_Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockICustomerService) S_Get(id int) (*models.Customer, error) {
	args := m.Called(id)
	if customer, ok := args.Get(0).(*models.Customer); ok {
		return customer, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockICustomerService) S_Auth(request dto.AuthRequest) (*dto.AuthResponse, error) {
	args := m.Called(request)
	if response, ok := args.Get(0).(*dto.AuthResponse); ok {
		return response, args.Error(1)
	}
	return nil, args.Error(1)
}
