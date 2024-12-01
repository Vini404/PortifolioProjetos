package services

import (
	"bytes"
	"errors"
	"io"
	"math/rand"
	"mime/multipart"
	"secbank.api/auth"
	dto "secbank.api/dto/customer"
	"secbank.api/interfaces/repository"
	"secbank.api/models"
	"strconv"
	"time"
)

type CustomerService struct {
	interfaces.ICustomerRepository
	interfaces.IAccountHolderRepository
	interfaces.IAccountRepository
	interfaces.IBalanceRepository
}

func (service *CustomerService) S_List() (*[]models.Customer, error) {
	return service.ICustomerRepository.R_List()
}

func (service *CustomerService) S_Create(customer models.Customer, file multipart.File) error {
	if err := service.checkCustomerExistence(customer.Email); err != nil {
		return err
	}

	customerID, err := service.ICustomerRepository.R_Create(customer)
	if err != nil {
		return err
	}

	if err := service.createAccountAndBalance(customerID); err != nil {
		return err
	}

	return service.processFacialRecognition(customerID, file)
}

func (service *CustomerService) S_Update(customer models.Customer) error {
	existingCustomer, err := service.ICustomerRepository.R_Get(customer.ID)
	if err != nil {
		return err
	}

	existingCustomer.Phone = customer.Phone
	existingCustomer.Birthday = customer.Birthday
	existingCustomer.Email = customer.Email

	return service.ICustomerRepository.R_Update(*existingCustomer)
}

func (service *CustomerService) S_Delete(id int) error {
	return service.ICustomerRepository.R_Delete(id)
}

func (service *CustomerService) S_Get(id int) (*models.Customer, error) {
	return service.ICustomerRepository.R_Get(id)
}

func (service *CustomerService) S_Auth(request dto.AuthRequest) (*dto.AuthResponse, error) {
	request.Validate()

	customer, err := service.ICustomerRepository.R_Get_By_Email(request.Email)
	if err != nil || customer == nil || customer.Password != request.Password {
		return nil, errors.New("Usuário ou senha incorreta.")
	}

	token, err := auth.GenerateJWT(customer.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{Token: token}, nil
}

// Helper methods for modularization
func (service *CustomerService) checkCustomerExistence(email string) error {
	customer, err := service.ICustomerRepository.R_Get_By_Email(email)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}
	if customer != nil {
		return errors.New("Já existe um usuário com o email informado")
	}
	return nil
}

func (service *CustomerService) createAccountAndBalance(customerID int) error {
	accountHolder := models.AccountHolder{
		IsActive:         true,
		CreatedTimeStamp: time.Now(),
		IDCustomer:       customerID,
	}

	holderID, err := service.IAccountHolderRepository.R_Create(accountHolder)
	if err != nil {
		return err
	}

	account := models.Account{
		CreatedTimeStamp: time.Now(),
		IsActive:         true,
		IDAccountHolder:  holderID,
		Number:           strconv.Itoa(generate7DigitNumber()),
		Digit:            truncateToOneDigit(holderID),
	}

	accountID, err := service.IAccountRepository.R_Create(account)
	if err != nil {
		return err
	}

	balance := models.Balance{
		IDAccount:        accountID,
		Amount:           200,
		AmountBlocked:    0,
		UpdatedTimeStamp: time.Now(),
	}

	_, err = service.IBalanceRepository.R_Create(balance)
	return err
}

func (service *CustomerService) processFacialRecognition(customerID int, file multipart.File) error {
	rekognitionService := NewRekognitionService("us-east-1")
	collectionID := "b7cff507-7306-4c37-a461-0ed736b7cdc5"

	if err := rekognitionService.CreateUser(collectionID, customerID); err != nil {
		return err
	}

	imageBytes, err := getFileBytes(file)
	if err != nil {
		return err
	}

	indexFaces, err := rekognitionService.IndexFaces(collectionID, imageBytes)
	if err != nil {
		return err
	}

	faceIDs := rekognitionService.GetFacesIDs(indexFaces)
	return rekognitionService.AssociateFacesToUser(collectionID, customerID, faceIDs)
}

func generate7DigitNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000000) + 1000000
}

func getFileBytes(file multipart.File) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	return buf.Bytes(), err
}

func truncateToOneDigit(holderID int) string {
	idStr := strconv.Itoa(holderID)

	if len(idStr) > 1 {
		return idStr[:1]
	}
	return idStr
}
