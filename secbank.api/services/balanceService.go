package services

import (
	dto "secbank.api/dto/balance"
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

func (service *BalanceService) S_Extract(accountID int) (*dto.ExtractResponse, error) {
	extract, err := service.IBalanceRepository.R_Extract(accountID)

	if err != nil {
		return nil, err
	}

	return extract, nil
}
