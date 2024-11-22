package controllers

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"secbank.api/controllers"
	mock_interfaces "secbank.api/mocks"
	"secbank.api/models"
	"testing"
)

func TestAccountController_List(t *testing.T) {
	mockService := new(mock_interfaces.MockIAccountService)
	controller := &controllers.AccountController{IAccountService: mockService}

	expectedAccounts := &[]models.Account{
		{ID: 1, Number: "123", Digit: "4", IsActive: true},
		{ID: 2, Number: "456", Digit: "7", IsActive: false},
	}

	mockService.On("S_List").Return(expectedAccounts, nil)

	req := httptest.NewRequest(http.MethodGet, "/accounts", nil)
	rec := httptest.NewRecorder()

	controller.List(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "123")
	assert.Contains(t, rec.Body.String(), "456")
}

func TestAccountController_Get(t *testing.T) {
	mockService := new(mock_interfaces.MockIAccountService)
	controller := &controllers.AccountController{IAccountService: mockService}

	expectedAccount := &models.Account{ID: 1, Number: "123", Digit: "4", IsActive: true}

	// Cenário de sucesso
	mockService.On("S_Get", 1).Return(expectedAccount, nil)

	req := httptest.NewRequest(http.MethodGet, "/accounts/1", nil)
	req = addChiURLParam(req, "id", "1")
	rec := httptest.NewRecorder()

	controller.Get(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "123")

}

func TestAccountController_Delete(t *testing.T) {
	mockService := new(mock_interfaces.MockIAccountService)
	controller := &controllers.AccountController{IAccountService: mockService}

	// Cenário de sucesso
	mockService.On("S_Delete", 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/accounts/1", nil)
	req = addChiURLParam(req, "id", "1")
	rec := httptest.NewRecorder()

	controller.Delete(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

// Helper function to add URL parameters
func addChiURLParam(req *http.Request, key, value string) *http.Request {
	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add(key, value)
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
	return req.WithContext(ctx)
}
