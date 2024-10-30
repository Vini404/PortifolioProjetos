package services

import (
	"fmt"
	dto "secbank.api/dto/transaction"
	interfaces "secbank.api/interfaces/repository"
	"secbank.api/models"
	"time"
)

type TransactionService struct {
	interfaces.IAccountRepository
	interfaces.IBalanceRepository
	interfaces.ITransactionRepository
}

func (service *TransactionService) Transfer(transferRequest dto.TransferRequest) error {
	creditAccount, errCreditAccountInformation := service.IAccountRepository.R_Get(transferRequest.IDCreditAccount)

	if errCreditAccountInformation != nil {
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
	}

	_, errInsertBalance := service.ITransactionRepository.R_Create(transaction)

	if errInsertBalance != nil {
		return errInsertBalance
	}

	return nil
}
