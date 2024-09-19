package services

import (
	"fmt"
	"secbank.api/internal/interfaces"
	"secbank.api/internal/models"
)

type CustomerService struct {
	interfaces.ICustomerService
}

func (service *CustomerService) List() ([]models.Customer, error) {
	allCustomers, err := service.List()
	if err != nil {
		fmt.Println(err.Error())
	}
	return allCustomers, nil
}
