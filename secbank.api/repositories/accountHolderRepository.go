package repositories

import (
	"secbank.api/interfaces"
	"secbank.api/models"
	"secbank.api/utils"
	"time"
)

type AccountHolderRepository struct {
	interfaces.IDbHandler
}

func (repository *AccountHolderRepository) R_List() (*[]models.AccountHolder, error) {
	var accountHolder []models.AccountHolder
	err := repository.Query("SELECT * FROM AccountHolder", &accountHolder)

	if err != nil {
		return nil, err
	}

	// Loop through rows, using Scan to assign column data to struct fields.

	return &accountHolder, nil
}

func (repository *AccountHolderRepository) R_Create(accountHolder models.AccountHolder) error {
	accountHolder.CreatedTimeStamp = time.Now()
	err := repository.Insert(accountHolder, "accountHolder")

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountHolderRepository) R_Update(accountHolder models.AccountHolder) error {
	accountHolder.UpdatedTimeStamp = time.Now()
	accountHolderToUpdate := utils.StructToMap(accountHolder)
	err := repository.Update(accountHolder.ID, "accountHolder", accountHolderToUpdate)

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountHolderRepository) R_Delete(id int) error {

	err := repository.Delete(id, "accountHolder")

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountHolderRepository) R_Get(id int) (*models.AccountHolder, error) {
	account := models.AccountHolder{}
	err := repository.Get(id, "accountHolder", &account)

	if err != nil {
		return nil, err
	}
	return &account, nil
}
