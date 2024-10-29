package interfaces

import (
	dto "secbank.api/dto/account"
	"secbank.api/models"
)

type IAccountService interface {
	S_List() (*[]models.Account, error)
	S_Create(customer models.Account) error
	S_Delete(id int) error
	S_Update(customer models.Account) error
	S_Get(id int) (*models.Account, error)
	S_GetInformationAccount(id int) (*dto.InformationAccountResponse, error)
}
