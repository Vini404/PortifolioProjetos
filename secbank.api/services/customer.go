package services

import (
	"errors"
	"fmt"
	"secbank.api/auth"
	dto "secbank.api/dto/customer"
	"secbank.api/interfaces/repository"
	"secbank.api/models"
)

type CustomerService struct {
	interfaces.ICustomerRepository
}

func (service *CustomerService) S_List() (*[]models.Customer, error) {
	allCustomers, err := service.R_List()
	return allCustomers, err
}

func (service *CustomerService) S_Create(customer models.Customer) error {
	err := service.ICustomerRepository.R_Create(customer)

	if err != nil {
		return err
	}

	return nil
}

func (service *CustomerService) S_Update(customer models.Customer) error {
	err := service.ICustomerRepository.R_Update(customer)

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
	customer, err := service.ICustomerRepository.R_Get_By_Email(request.Email)

	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.New("Usuario não encontrado")
	}

	if customer.Password != request.Password {
		return nil, errors.New("Usuario ou senha incorreta.")
	}

	token, err := auth.GenerateJWT(customer.Email)

	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{Token: token}, nil
}
