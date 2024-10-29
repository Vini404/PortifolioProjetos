package repositories

import (
	"secbank.api/interfaces"
	"secbank.api/models"
)

type BalanceRepository struct {
	interfaces.IDbHandler
}

func (repository *BalanceRepository) R_GetByAccountID(accountID int) (*models.Balance, error) {
	balance := models.Balance{}
	err := repository.QueryWithParamSingleRow("SELECT * FROM Balance WHERE idaccount=$1", &balance, accountID)

	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (repository *BalanceRepository) R_Create(balance models.Balance) (int, error) {
	id, err := repository.Insert(balance, "balance")
	if err != nil {
		return 0, err
	}

	return id, nil
}
