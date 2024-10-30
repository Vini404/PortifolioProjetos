package dto

type TransferUserRequest struct {
	IDCreditAccount int     `json:"credit_account"`
	Amount          float64 `json:"amount"`
}
