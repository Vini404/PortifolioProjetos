package services

import (
	"fmt"
	"secbank.api/interfaces"
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
