package services

import (
	"errors"
	"mime/multipart"
	dto "secbank.api/dto/transaction"
	interfaces "secbank.api/interfaces/repository"
	"secbank.api/models"
	"strconv"
	"time"
)

type TransactionService struct {
	interfaces.IAccountRepository
	interfaces.IBalanceRepository
	interfaces.ITransactionRepository
	interfaces.ICustomerRepository
}

func (service *TransactionService) Transfer(transferRequest dto.TransferRequest, file multipart.File) error {
	creditAccount, err := service.getCreditAccount(transferRequest.NumberCreditAccount, transferRequest.DigitCreditAccount)
	if err != nil {
		return err
	}

	debitAccount, err := service.IAccountRepository.R_GetAccountByCustomer(transferRequest.IDCustomerOriginAccount)
	if err != nil {
		return err
	}

	if err := service.validateDebitAccountBalance(debitAccount.ID, transferRequest.Amount); err != nil {
		return err
	}

	if err := service.updateAccountBalances(debitAccount.ID, creditAccount.ID, transferRequest.Amount); err != nil {
		return err
	}

	if err := service.recordTransaction(debitAccount.ID, creditAccount.ID, transferRequest.Amount); err != nil {
		return err
	}

	return service.validateFacialRecognition(file)
}

// Helper methods
func (service *TransactionService) getCreditAccount(accountNumber, accountDigit string) (*models.Account, error) {
	account, err := service.IAccountRepository.R_Get_By_Number_And_Digit(accountNumber, accountDigit)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("A conta informada não existe.")
		}
		return nil, err
	}
	return account, nil
}

func (service *TransactionService) validateDebitAccountBalance(accountID int, amount float64) error {
	balance, err := service.IBalanceRepository.R_GetByAccountID(accountID)
	if err != nil {
		return err
	}

	unlockedAmount := balance.Amount - balance.AmountBlocked
	if unlockedAmount < amount {
		return errors.New("Você não possui saldo disponível para realizar essa transferência.")
	}
	return nil
}

func (service *TransactionService) updateAccountBalances(debitAccountID, creditAccountID int, amount float64) error {
	// Update debit account balance
	debitBalance, err := service.IBalanceRepository.R_GetByAccountID(debitAccountID)
	if err != nil {
		return err
	}
	debitBalance.Amount -= amount

	if err := service.IBalanceRepository.R_Update(debitBalance); err != nil {
		return err
	}

	// Update credit account balance
	creditBalance, err := service.IBalanceRepository.R_GetByAccountID(creditAccountID)
	if err != nil {
		return err
	}
	creditBalance.Amount += amount

	return service.IBalanceRepository.R_Update(creditBalance)
}

func (service *TransactionService) recordTransaction(debitAccountID, creditAccountID int, amount float64) error {
	transaction := models.Transaction{
		CreatedTimeStamp: time.Now(),
		IDCreditAccount:  creditAccountID,
		IDDebitAccount:   debitAccountID,
		Amount:           amount,
		Description:      "Transferência entre contas",
		TransactionType:  1,
	}

	_, err := service.ITransactionRepository.R_Create(transaction)
	return err
}

func (service *TransactionService) validateFacialRecognition(file multipart.File) error {
	imageBytes, err := getFileBytes(file)
	if err != nil {
		return err
	}

	collectionID := "b7cff507-7306-4c37-a461-0ed736b7cdc5"
	rekognitionService := NewRekognitionService("us-east-1")

	users, err := rekognitionService.SearchUsersByImage(collectionID, imageBytes)
	if err != nil || len(users.UserMatches) == 0 {
		return errors.New("Falha na validação de reconhecimento facial.")
	}

	userID, err := strconv.Atoi(*users.UserMatches[0].User.UserId)
	if err != nil {
		return err
	}

	_, err = service.ICustomerRepository.R_Get(userID)
	if err != nil {
		return errors.New("Falha na validação de reconhecimento facial.")
	}

	return nil
}
