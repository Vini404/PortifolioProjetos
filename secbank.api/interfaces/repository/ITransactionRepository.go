package interfaces

import (
	"secbank.api/models"
)

type ITransactionRepository interface {
	R_Create(balance models.Transaction) (int, error)
}
