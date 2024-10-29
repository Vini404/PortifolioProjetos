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
	err := repository.Query("SELECT * FROM accountholder", &accountHolder)

	if err != nil {
		return nil, err
	}

	// Loop through rows, using Scan to assign column data to struct fields.

	return &accountHolder, nil
}

func (repository *AccountHolderRepository) R_Create(accountHolder models.AccountHolder) (int, error) {
	accountHolder.CreatedTimeStamp = time.Now()
	id, err := repository.Insert(accountHolder, "accountholder")

	if err != nil {
		return 0, err
	}
	return id, err
}

func (repository *AccountHolderRepository) R_Update(accountHolder models.AccountHolder) error {
	accountHolder.UpdatedTimeStamp = time.Now()
	accountHolderToUpdate := utils.StructToMap(accountHolder)
	err := repository.Update(accountHolder.ID, "accountholder", accountHolderToUpdate)

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountHolderRepository) R_Delete(id int) error {

	err := repository.Delete(id, "accountholder")

	if err != nil {
		return err
	}
	return nil
}

func (repository *AccountHolderRepository) R_Get(id int) (*models.AccountHolder, error) {
	account := models.AccountHolder{}
	err := repository.Get(id, "accountholder", &account)

	if err != nil {
		return nil, err
	}
	return &account, nil
}
