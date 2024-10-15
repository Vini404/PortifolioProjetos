package repositories

import (
	"secbank.api/interfaces"
	"secbank.api/models"
	"secbank.api/utils"
	"time"
)

type AccountRepository struct {
	interfaces.IDbHandler
}

func (repository *AccountRepository) R_List() (*[]models.Account, error) {
	var account []models.Account
	err := repository.Query("SELECT * FROM Account", &account)

	if err != nil {
		return nil, err
	}

	// Loop through rows, using Scan to assign column data to struct fields.

	return &account, nil
}

func (repository *AccountRepository) R_Create(account models.Account) error {
	account.CreatedTimeStamp = time.Now()
	err := repository.Insert(account, "account")

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountRepository) R_Update(account models.Account) error {
	account.UpdatedTimeStamp = time.Now()
	accountToUpdate := utils.StructToMap(account)
	err := repository.Update(account.ID, "customer", accountToUpdate)

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountRepository) R_Delete(id int) error {

	err := repository.Delete(id, "account")

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountRepository) R_Get(id int) (*models.Account, error) {
	account := models.Account{}
	err := repository.Get(id, "account", &account)

	if err != nil {
		return nil, err
	}
	return &account, nil
}
