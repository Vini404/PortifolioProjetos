// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\VINICIUS-VN\Desktop\projetos\PortifolioProjetos\secbank.api\interfaces\repository/ICustomerRepository.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "secbank.api/models"
)

// MockICustomerRepository is a mock of ICustomerRepository interface.
type MockICustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICustomerRepositoryMockRecorder
}

// MockICustomerRepositoryMockRecorder is the mock recorder for MockICustomerRepository.
type MockICustomerRepositoryMockRecorder struct {
	mock *MockICustomerRepository
}

// NewMockICustomerRepository creates a new mock instance.
func NewMockICustomerRepository(ctrl *gomock.Controller) *MockICustomerRepository {
	mock := &MockICustomerRepository{ctrl: ctrl}
	mock.recorder = &MockICustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICustomerRepository) EXPECT() *MockICustomerRepositoryMockRecorder {
	return m.recorder
}

// R_Create mocks base method.
func (m *MockICustomerRepository) R_Create(customer models.Customer) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R_Create", customer)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// R_Create indicates an expected call of R_Create.
func (mr *MockICustomerRepositoryMockRecorder) R_Create(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R_Create", reflect.TypeOf((*MockICustomerRepository)(nil).R_Create), customer)
}

// R_Delete mocks base method.
func (m *MockICustomerRepository) R_Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R_Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// R_Delete indicates an expected call of R_Delete.
func (mr *MockICustomerRepositoryMockRecorder) R_Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R_Delete", reflect.TypeOf((*MockICustomerRepository)(nil).R_Delete), id)
}

// R_Get mocks base method.
func (m *MockICustomerRepository) R_Get(id int) (*models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R_Get", id)
	ret0, _ := ret[0].(*models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// R_Get indicates an expected call of R_Get.
func (mr *MockICustomerRepositoryMockRecorder) R_Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R_Get", reflect.TypeOf((*MockICustomerRepository)(nil).R_Get), id)
}

// R_Get_By_Email mocks base method.
func (m *MockICustomerRepository) R_Get_By_Email(email string) (*models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R_Get_By_Email", email)
	ret0, _ := ret[0].(*models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// R_Get_By_Email indicates an expected call of R_Get_By_Email.
func (mr *MockICustomerRepositoryMockRecorder) R_Get_By_Email(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R_Get_By_Email", reflect.TypeOf((*MockICustomerRepository)(nil).R_Get_By_Email), email)
}

// R_List mocks base method.
func (m *MockICustomerRepository) R_List() (*[]models.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R_List")
	ret0, _ := ret[0].(*[]models.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// R_List indicates an expected call of R_List.
func (mr *MockICustomerRepositoryMockRecorder) R_List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R_List", reflect.TypeOf((*MockICustomerRepository)(nil).R_List))
}

// R_Update mocks base method.
func (m *MockICustomerRepository) R_Update(customer models.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R_Update", customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// R_Update indicates an expected call of R_Update.
func (mr *MockICustomerRepositoryMockRecorder) R_Update(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R_Update", reflect.TypeOf((*MockICustomerRepository)(nil).R_Update), customer)
}
