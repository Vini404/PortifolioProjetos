package services

import (
	"bytes"
	"errors"
	"fmt"
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
	allCustomers, err := service.ICustomerRepository.R_List()
	return allCustomers, err
}

func (service *CustomerService) S_Create(customer models.Customer, file multipart.File) error {

	id, err := service.ICustomerRepository.R_Create(customer)

	if err != nil {
		return err
	}

	accountHolder := models.AccountHolder{
		IsActive:         true,
		CreatedTimeStamp: time.Now(),
		IDCustomer:       id,
	}

	idAccountHolder, errAccountHolder := service.IAccountHolderRepository.R_Create(accountHolder)

	if errAccountHolder != nil {
		return errAccountHolder
	}

	account := models.Account{

		CreatedTimeStamp: time.Now(),
		IsActive:         true,
		IDAccountHolder:  idAccountHolder,
		Number:           strconv.Itoa(generate7DigitNumber()),
		Digit:            strconv.Itoa(idAccountHolder),
	}

	accountID, errAccount := service.IAccountRepository.R_Create(account)

	if errAccount != nil {
		return errAccount
	}

	balance := models.Balance{
		IDAccount:        accountID,
		Amount:           200,
		AmountBlocked:    0,
		UpdatedTimeStamp: time.Now(),
	}

	_, errBalance := service.IBalanceRepository.R_Create(balance)

	if errBalance != nil {
		return err
	}

	rekognitionService := NewRekognitionService("us-east-1")

	collectionID := "b7cff507-7306-4c37-a461-0ed736b7cdc5"

	errCreateUserRekognition := rekognitionService.CreateUser(collectionID, id)

	if errCreateUserRekognition != nil {
		return errCreateUserRekognition
	}

	imageBytes, errGetImageBytes := getFileBytes(file)

	if errGetImageBytes != nil {
		return errGetImageBytes
	}

	indexFaces, errIndexFaces := rekognitionService.IndexFaces(collectionID, imageBytes)

	if errIndexFaces != nil {
		return errIndexFaces
	}

	facesIds := rekognitionService.GetFacesIDs(indexFaces)

	errAssociateFacesToUser := rekognitionService.AssociateFacesToUser(collectionID, id, facesIds)

	if errAssociateFacesToUser != nil {
		return errAssociateFacesToUser
	}

	return nil
}

func (service *CustomerService) S_Update(customer models.Customer) error {
	customerOriginal, _ := service.ICustomerRepository.R_Get(customer.ID)

	customerOriginal.Phone = customer.Phone
	customerOriginal.Birthday = customer.Birthday
	customerOriginal.Email = customer.Email

	err := service.ICustomerRepository.R_Update(*customerOriginal)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (service *CustomerService) S_Delete(id int) error {
	err := service.ICustomerRepository.R_Delete(id)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (service *CustomerService) S_Get(id int) (*models.Customer, error) {
	customer, err := service.ICustomerRepository.R_Get(id)

	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (service *CustomerService) S_Auth(request dto.AuthRequest) (*dto.AuthResponse, error) {
	request.Validate()

	customer, err := service.ICustomerRepository.R_Get_By_Email(request.Email)

	if err != nil {

		if err.Error() != "sql: no rows in result set" {
			return nil, fmt.Errorf("Usuario ou senha incorreta.")
		}

		return nil, err
	}

	if customer == nil {
		return nil, errors.New("Usuario ou senha incorreta.")
	}

	if customer.Password != request.Password {
		return nil, errors.New("Usuario ou senha incorreta.")
	}

	token, err := auth.GenerateJWT(customer.ID)

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{Token: token}, nil
}

func generate7DigitNumber() int {
	rand.Seed(time.Now().UnixNano())    // Seed the random number generator
	return rand.Intn(9000000) + 1000000 // Generates a number between 1,000,000 and 9,999,999
}

func getFileBytes(file multipart.File) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
