package interfaces

import "secbank.api/internal/models"

type ICustomerService interface {
	List() ([]models.Customer, error)
}
