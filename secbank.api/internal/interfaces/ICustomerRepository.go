package interfaces

import "secbank.api/internal/models"

type ICustomerRepository interface {
	List(name string) (models.Customer, error)
}
