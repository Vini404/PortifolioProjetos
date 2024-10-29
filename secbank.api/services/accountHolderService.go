package services

import (
	"fmt"
	"secbank.api/interfaces/repository"
	"secbank.api/models"
)

type AccountHolderService struct {
	interfaces.IAccountHolderRepository
}

func (service *AccountHolderService) S_List() (*[]models.AccountHolder, error) {
	allAccountHolders, err := service.R_List()
	return allAccountHolders, err
}

func (service *AccountHolderService) S_Create(accountHolder models.AccountHolder) (int, error) {
	id, err := service.IAccountHolderRepository.R_Create(accountHolder)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *AccountHolderService) S_Update(accountHolder models.AccountHolder) error {
	err := service.IAccountHolderRepository.R_Update(accountHolder)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (service *AccountHolderService) S_Delete(id int) error {
	err := service.IAccountHolderRepository.R_Delete(id)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (service *AccountHolderService) S_Get(id int) (*models.AccountHolder, error) {
	accountHolder, err := service.IAccountHolderRepository.R_Get(id)

	if err != nil {
		return nil, err
	}
	return accountHolder, nil
}
