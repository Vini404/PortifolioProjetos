package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"secbank.api/controllers"
	dto "secbank.api/dto/customer"
	mock_interfaces "secbank.api/mocks"
	"secbank.api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerController_List(t *testing.T) {
	mockService := new(mock_interfaces.MockICustomerService)
	controller := &controllers.CustomerController{ICustomerService: mockService}

	expectedCustomers := &[]models.Customer{
		{ID: 1, FullName: "John Doe", Email: "john@example.com"},
		{ID: 2, FullName: "Jane Doe", Email: "jane@example.com"},
	}

	mockService.On("S_List").Return(expectedCustomers, nil)

	req := httptest.NewRequest(http.MethodGet, "/customers", nil)
	rec := httptest.NewRecorder()

	controller.List(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])
}

func TestCustomerController_Auth(t *testing.T) {
	mockService := new(mock_interfaces.MockICustomerService)
	controller := &controllers.CustomerController{ICustomerService: mockService}

	authRequest := dto.AuthRequest{Email: "john@example.com", Password: "password123"}
	authResponse := &dto.AuthResponse{Token: "mock-token"}

	mockService.On("S_Auth", authRequest).Return(authResponse, nil)

	body, _ := json.Marshal(authRequest)
	req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	controller.Auth(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.True(t, response["success"].(bool))
	assert.Equal(t, "mock-token", response["data"].(map[string]interface{})["token"])
}

func TestCustomerController_Get(t *testing.T) {
	mockService := new(mock_interfaces.MockICustomerService)
	controller := &controllers.CustomerController{ICustomerService: mockService}

	expectedCustomer := &models.Customer{ID: 1, FullName: "John Doe", Email: "john@example.com"}
	mockService.On("S_Get", 1).Return(expectedCustomer, nil)

	req := httptest.NewRequest(http.MethodGet, "/customers/1", nil)
	req = addChiURLParam(req, "id", "1")
	rec := httptest.NewRecorder()

	controller.Get(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])
}
