package interfaces

import (
	"secbank.api/models"
)

type IAccountHolderService interface {
	S_List() (*[]models.AccountHolder, error)
	S_Create(customer models.AccountHolder) (int, error)
	S_Delete(id int) error
	S_Update(customer models.AccountHolder) error
	S_Get(id int) (*models.AccountHolder, error)
}
