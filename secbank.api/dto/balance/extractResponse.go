package dto

type ExtractResponse struct {
	OperationName string  `json:"operation_name"`
	Amount        float64 `json:"amount"`
	TransferType  string  `json:"transfer_type"`
}
