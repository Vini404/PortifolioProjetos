package interfaces

import (
	dto "secbank.api/dto/account"
	"secbank.api/models"
)

type IAccountRepository interface {
	R_List() (*[]models.Account, error)
	R_Create(customer models.Account) (int, error)
	R_Update(customer models.Account) error
	R_Delete(id int) error
	R_Get(id int) (*models.Account, error)
	R_GetInformationAccount(id int) (*dto.InformationAccountResponse, error)
}
