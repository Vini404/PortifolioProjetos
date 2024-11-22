package test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	dto "secbank.api/dto/customer"
	mock_interfaces "secbank.api/mocks"
	"secbank.api/models"
	"secbank.api/services"
	"testing"
)

func TestCustomerService_S_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockICustomerRepository(ctrl)
	service := &services.CustomerService{ICustomerRepository: mockRepo}

	expectedCustomers := &[]models.Customer{
		{ID: 1, FullName: "John Doe", Email: "john@example.com"},
		{ID: 2, FullName: "Jane Doe", Email: "jane@example.com"},
	}

	mockRepo.EXPECT().R_List().Return(expectedCustomers, nil)

	customers, err := service.S_List()
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)
}

func TestCustomerService_S_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockICustomerRepository(ctrl)
	service := &services.CustomerService{ICustomerRepository: mockRepo}

	customer := models.Customer{ID: 1, Phone: "123456789", Email: "newemail@example.com"}
	originalCustomer := &models.Customer{ID: 1, FullName: "John Doe", Phone: "987654321", Email: "oldemail@example.com"}

	mockRepo.EXPECT().R_Get(1).Return(originalCustomer, nil)
	mockRepo.EXPECT().R_Update(gomock.Any()).Return(nil)

	err := service.S_Update(customer)
	assert.NoError(t, err)
}

func TestCustomerService_S_Auth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockICustomerRepository(ctrl)
	service := &services.CustomerService{ICustomerRepository: mockRepo}

	request := dto.AuthRequest{Email: "user@example.com", Password: "password123"}
	customer := &models.Customer{ID: 1, Email: "user@example.com", Password: "password123"}

	mockRepo.EXPECT().R_Get_By_Email("user@example.com").Return(customer, nil)

	authResponse, err := service.S_Auth(request)
	assert.NoError(t, err)
	assert.NotNil(t, authResponse)

	// Testando senha incorreta
	customer.Password = "wrongpassword"
	mockRepo.EXPECT().R_Get_By_Email("user@example.com").Return(customer, nil)

	authResponse, err = service.S_Auth(request)
	assert.Error(t, err)
	assert.Nil(t, authResponse)

	// Testando email inexistente
	mockRepo.EXPECT().R_Get_By_Email("user@example.com").Return(nil, errors.New("sql: no rows in result set"))

	authResponse, err = service.S_Auth(request)
	assert.Error(t, err)
	assert.Nil(t, authResponse)
}
