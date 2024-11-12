package services

import (
	"fmt"
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
	creditAccount, errCreditAccountInformation := service.IAccountRepository.R_Get_By_Number_And_Digit(transferRequest.NumberCreditAccount, transferRequest.DigitCreditAccount)

	if errCreditAccountInformation != nil {

		if errCreditAccountInformation.Error() == "sql: no rows in result set" {
			return fmt.Errorf("A conta informada não existe.")
		}

		return errCreditAccountInformation
	}
	debitAccount, errDebitAccountInformation := service.IAccountRepository.R_GetAccountByCustomer(transferRequest.IDCustomerOriginAccount)

	if errDebitAccountInformation != nil {
		return errDebitAccountInformation
	}

	balanceDebitAccount, errGetBalanceDebitAccount := service.IBalanceRepository.R_GetByAccountID(debitAccount.ID)

	if errGetBalanceDebitAccount != nil {
		return errGetBalanceDebitAccount
	}

	unlockedAmount := balanceDebitAccount.Amount - balanceDebitAccount.AmountBlocked

	if unlockedAmount < transferRequest.Amount {
		return fmt.Errorf("Você não possui saldo disponivel para realizar essa transferencia")
	}

	balanceCreditAccount, errGetBalanceCreditAccount := service.IBalanceRepository.R_GetByAccountID(creditAccount.ID)

	if errGetBalanceCreditAccount != nil {
		return errGetBalanceCreditAccount
	}

	balanceCreditAccount.Amount += transferRequest.Amount
	balanceDebitAccount.Amount -= transferRequest.Amount

	errUpdateBalanceCreditAccount := service.IBalanceRepository.R_Update(balanceCreditAccount)

	if errUpdateBalanceCreditAccount != nil {
		return errUpdateBalanceCreditAccount
	}

	errUpdateBalanceDebitAccount := service.IBalanceRepository.R_Update(balanceDebitAccount)

	if errUpdateBalanceDebitAccount != nil {
		return errUpdateBalanceDebitAccount
	}

	transaction := models.Transaction{
		CreatedTimeStamp: time.Now(),
		IDCreditAccount:  creditAccount.ID,
		IDDebitAccount:   debitAccount.ID,
		Amount:           transferRequest.Amount,
		Description:      "Transferencia entre contas",
		TransactionType:  1,
	}

	_, errInsertBalance := service.ITransactionRepository.R_Create(transaction)

	if errInsertBalance != nil {
		return errInsertBalance
	}

	imageBytes, errGetImageBytes := getFileBytes(file)

	if errGetImageBytes != nil {
		return errGetImageBytes
	}

	collectionID := "b7cff507-7306-4c37-a461-0ed736b7cdc5"

	rekognitionService := NewRekognitionService("us-east-1")

	users, errorSearchUsers := rekognitionService.SearchUsersByImage(collectionID, imageBytes)

	if errorSearchUsers != nil {
		return errorSearchUsers
	}

	if len(users.UserMatches) > 0 {
		userID, errParseUserID := strconv.Atoi(*users.UserMatches[0].User.UserId)

		if errParseUserID != nil {
			return errParseUserID
		}

		_, err := service.ICustomerRepository.R_Get(userID)

		if err != nil {
			return fmt.Errorf("Falha na validação de reconhecimento facial.")
		}

	} else {
		return fmt.Errorf("Falha na validação de reconhecimento facial.")
	}

	return nil
}
