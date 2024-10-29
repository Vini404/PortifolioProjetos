package services

import (
	interfaces "secbank.api/interfaces/repository"
	"secbank.api/models"
)

type BalanceService struct {
	interfaces.IBalanceRepository
}

func (service *BalanceService) S_GetByAccountID(accountID int) (*models.Balance, error) {
	balance, err := service.IBalanceRepository.R_GetByAccountID(accountID)

	if err != nil {
		return nil, err
	}
	
	return balance, nil
}
