package services

import (
	"fmt"
	"secbank.api/interfaces"
	"secbank.api/models"
)

type CustomerService struct {
	interfaces.ICustomerRepository
}

func (service *CustomerService) List() (*[]models.Customer, error) {
	allCustomers, err := service.ListAllCustomer()
	if err != nil {
		fmt.Println(err.Error())
	}
	return allCustomers, nil
}
