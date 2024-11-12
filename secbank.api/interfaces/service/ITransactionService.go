package interfaces

import (
	"mime/multipart"
	dto "secbank.api/dto/transaction"
)

type ITransactionService interface {
	Transfer(transferRequest dto.TransferRequest, file multipart.File) error
}
