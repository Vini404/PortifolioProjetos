package dto

type ExtractResponse struct {
	OperationName string  `db:"operation_name"`
	Amount        float64 `db:"amount"`
	TransferType  string  `db:"transfer_type"`
}
