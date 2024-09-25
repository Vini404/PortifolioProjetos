package interfaces

import "secbank.api/models"

type ICustomerService interface {
	List() (*[]models.Customer, error)
}
