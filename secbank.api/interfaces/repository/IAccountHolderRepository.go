package interfaces

import (
	"secbank.api/models"
)

type IAccountHolderRepository interface {
	R_List() (*[]models.AccountHolder, error)
	R_Create(customer models.AccountHolder) (int, error)
	R_Update(customer models.AccountHolder) error
	R_Delete(id int) error
	R_Get(id int) (*models.AccountHolder, error)
}
