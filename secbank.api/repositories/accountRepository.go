package repositories

import (
	dto "secbank.api/dto/account"
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

func (repository *AccountRepository) R_Create(account models.Account) (int, error) {
	account.CreatedTimeStamp = time.Now()
	id, err := repository.Insert(account, "account")

	if err != nil {
		return 0, err
	}
	return id, nil
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

func (repository *AccountRepository) R_GetInformationAccount(id int) (*dto.InformationAccountResponse, error) {
	informationAccount := dto.InformationAccountResponse{}
	sql := `
		SELECT 
		    concat(a.Number,'-',a.Digit) as AccountNumber, 
		    c.FullName as CustomerName,
		    c.ID as CustomerID,
		    a.ID as IDAccount
			from Account a
			inner join AccountHolder ah on ah.ID = a.IDAccountHolder
			inner join customer c on c.ID = ah.IdCustomer
			where a.ID = $1`
	err := repository.QueryWithParamSingleRow(sql, &informationAccount, id)

	if err != nil {
		return nil, err
	}
	return &informationAccount, nil
}

func (repository *AccountRepository) R_Get_By_Number_And_Digit(number string, digit string) (*models.Account, error) {
	account := models.Account{}
	sql := `
		SELECT 
		    *
			from Account a
			where a.Number = $1 and a.Digit = $2`
	err := repository.QueryWithParamSingleRow(sql, &account, number, digit)

	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (repository *AccountRepository) R_GetAccountByCustomer(customerID int) (*models.Account, error) {
	account := models.Account{}
	sql := `
		SELECT 
		   	a.*
			from Account a
			inner join AccountHolder ah on ah.ID = a.IDAccountHolder
			inner join customer c on c.ID = ah.IdCustomer
			where c.ID = $1`
	err := repository.QueryWithParamSingleRow(sql, &account, customerID)

	if err != nil {
		return nil, err
	}
	return &account, nil
}
