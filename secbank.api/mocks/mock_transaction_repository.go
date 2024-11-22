// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\VINICIUS-VN\Desktop\projetos\PortifolioProjetos\secbank.api\interfaces\repository/ITransactionRepository.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "secbank.api/models"
)

// MockITransactionRepository is a mock of ITransactionRepository interface.
type MockITransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionRepositoryMockRecorder
}

// MockITransactionRepositoryMockRecorder is the mock recorder for MockITransactionRepository.
type MockITransactionRepositoryMockRecorder struct {
	mock *MockITransactionRepository
}

// NewMockITransactionRepository creates a new mock instance.
func NewMockITransactionRepository(ctrl *gomock.Controller) *MockITransactionRepository {
	mock := &MockITransactionRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionRepository) EXPECT() *MockITransactionRepositoryMockRecorder {
	return m.recorder
}

// R_Create mocks base method.
func (m *MockITransactionRepository) R_Create(balance models.Transaction) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R_Create", balance)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// R_Create indicates an expected call of R_Create.
func (mr *MockITransactionRepositoryMockRecorder) R_Create(balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R_Create", reflect.TypeOf((*MockITransactionRepository)(nil).R_Create), balance)
}
