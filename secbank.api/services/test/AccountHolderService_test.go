package test

import (
	"reflect"
	mock_interfaces "secbank.api/mocks"
	"secbank.api/models"
	"secbank.api/services"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestAccountHolderService_S_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountHolderRepository(ctrl)
	mockService := services.AccountHolderService{IAccountHolderRepository: mockRepo}

	expectedAccountHolders := &[]models.AccountHolder{
		{ID: 1, IDCustomer: 101, IsActive: true, CreatedTimeStamp: time.Now()},
		{ID: 2, IDCustomer: 102, IsActive: false, CreatedTimeStamp: time.Now()},
	}

	mockRepo.EXPECT().R_List().Return(expectedAccountHolders, nil)

	result, err := mockService.S_List()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expectedAccountHolders) {
		t.Errorf("expected %v, got %v", expectedAccountHolders, result)
	}
}

func TestAccountHolderService_S_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountHolderRepository(ctrl)
	mockService := services.AccountHolderService{IAccountHolderRepository: mockRepo}

	accountHolder := models.AccountHolder{IDCustomer: 101, IsActive: true}
	expectedID := 1

	mockRepo.EXPECT().R_Create(accountHolder).Return(expectedID, nil)

	id, err := mockService.S_Create(accountHolder)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if id != expectedID {
		t.Errorf("expected ID %d, got %d", expectedID, id)
	}
}

func TestAccountHolderService_S_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountHolderRepository(ctrl)
	mockService := services.AccountHolderService{IAccountHolderRepository: mockRepo}

	accountHolder := models.AccountHolder{ID: 1, IDCustomer: 101, IsActive: false}

	mockRepo.EXPECT().R_Update(accountHolder).Return(nil)

	err := mockService.S_Update(accountHolder)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAccountHolderService_S_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountHolderRepository(ctrl)
	mockService := services.AccountHolderService{IAccountHolderRepository: mockRepo}

	id := 1

	mockRepo.EXPECT().R_Delete(id).Return(nil)

	err := mockService.S_Delete(id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAccountHolderService_S_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_interfaces.NewMockIAccountHolderRepository(ctrl)
	mockService := services.AccountHolderService{IAccountHolderRepository: mockRepo}

	expectedAccountHolder := &models.AccountHolder{
		ID:               1,
		IDCustomer:       101,
		IsActive:         true,
		CreatedTimeStamp: time.Now(),
	}

	mockRepo.EXPECT().R_Get(1).Return(expectedAccountHolder, nil)

	accountHolder, err := mockService.S_Get(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(accountHolder, expectedAccountHolder) {
		t.Errorf("expected %v, got %v", expectedAccountHolder, accountHolder)
	}
}
