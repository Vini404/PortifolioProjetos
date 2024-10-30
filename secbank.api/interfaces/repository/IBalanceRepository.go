package interfaces

import (
	dto "secbank.api/dto/balance"
	"secbank.api/models"
)

type IBalanceRepository interface {
	R_GetByAccountID(accountID int) (*models.Balance, error)
	R_Create(balance models.Balance) (int, error)
	R_Extract(accountID int) (*dto.ExtractResponse, error)
	R_Update(balance *models.Balance) error
}
