package repositories

import (
	"secbank.api/interfaces"
	"secbank.api/models"
	"time"
)

type TransactionRepository struct {
	interfaces.IDbHandler
}

func (repository *TransactionRepository) R_Create(transaction models.Transaction) (int, error) {
	transaction.CreatedTimeStamp = time.Now()
	id, err := repository.Insert(transaction, "transactions")
	if err != nil {
		return 0, err
	}

	return id, nil
}
