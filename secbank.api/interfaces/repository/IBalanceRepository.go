package interfaces

import "secbank.api/models"

type IBalanceRepository interface {
	R_GetByAccountID(accountID int) (*models.Balance, error)
	R_Create(balance models.Balance) (int, error)
}
