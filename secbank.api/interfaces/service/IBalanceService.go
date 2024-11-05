package interfaces

import (
	dto "secbank.api/dto/balance"
	"secbank.api/models"
)

type IBalanceService interface {
	S_GetByAccountID(accountID int) (*models.Balance, error)
	S_Extract(accountID int) ([]*dto.ExtractResponse, error)
}
