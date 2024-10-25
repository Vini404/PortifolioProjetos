package interfaces

import (
	"secbank.api/models"
)

type ICustomerRepository interface {
	R_List() (*[]models.Customer, error)
	R_Create(customer models.Customer) error
	R_Update(customer models.Customer) error
	R_Delete(id int) error
	R_Get(id int) (*models.Customer, error)
	R_Get_By_Email(email string) (*models.Customer, error)
}
