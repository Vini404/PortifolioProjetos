package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"secbank.api/controllers"
	"secbank.api/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetResponseError(t *testing.T) {
	rec := httptest.NewRecorder()

	// Erro simulado
	controllers.SetResponseError(rec, assert.AnError)

	// Verificar status HTTP
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Verificar cabeçalho Content-Type
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	// Verificar corpo da resposta
	var response dto.Response
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.False(t, response.Success)
	assert.Equal(t, "assert.AnError general error for testing", response.MessageError)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestSetResponseSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	// Chamar o método
	controllers.SetResponseSuccess(rec)

	// Verificar status HTTP
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verificar cabeçalho Content-Type
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	// Verificar corpo da resposta
	var response dto.Response
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.True(t, response.Success)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.WithinDuration(t, time.Now(), response.Timestamp, time.Second)
}

func TestSetResponseSuccessWithPayload(t *testing.T) {
	rec := httptest.NewRecorder()

	// Dados de payload simulados
	payload := map[string]string{"key": "value"}

	// Chamar o método
	controllers.SetResponseSuccessWithPayload(rec, payload)

	// Verificar status HTTP
	assert.Equal(t, http.StatusOK, rec.Code)

	// Verificar cabeçalho Content-Type
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	// Verificar corpo da resposta
	var response dto.Response
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.True(t, response.Success)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Comparar o payload como map[string]interface{} para evitar incompatibilidade de tipos
	expectedPayload := map[string]interface{}{
		"key": "value",
	}
	assert.Equal(t, expectedPayload, response.Data)
	assert.WithinDuration(t, time.Now(), response.Timestamp, time.Second)
}
