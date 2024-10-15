package services

import (
	"fmt"
	"secbank.api/interfaces/repository"
	"secbank.api/models"
)

type AccountService struct {
	interfaces.IAccountRepository
}

func (service *AccountService) S_List() (*[]models.Account, error) {
	allAccounts, err := service.R_List()
	return allAccounts, err
}

func (service *AccountService) S_Create(account models.Account) error {
	err := service.IAccountRepository.R_Create(account)

	if err != nil {
		return err
	}

	return nil
}

func (service *AccountService) S_Update(account models.Account) error {
	err := service.IAccountRepository.R_Update(account)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (service *AccountService) S_Delete(id int) error {
	err := service.IAccountRepository.R_Delete(id)

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (service *AccountService) S_Get(id int) (*models.Account, error) {
	account, err := service.IAccountRepository.R_Get(id)

	if err != nil {
		return nil, err
	}
	return account, nil
}
