package interfaces

import "secbank.api/models"

type ICustomerRepository interface {
	ListAllCustomer() (*[]models.Customer, error)
}
