package interfaces

import dto "secbank.api/dto/transaction"

type ITransactionService interface {
	Transfer(transferRequest dto.TransferRequest) error
}
