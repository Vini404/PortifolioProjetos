package interfaces

import (
	"secbank.api/models"
)

type IBalanceService interface {
	S_GetByAccountID(accountID int) (*models.Balance, error)
}
