package dto

type TransferUserRequest struct {
	DigitCreditAccount  string  `json:"digit_credit_account"`
	NumberCreditAccount string  `json:"number_credit_account"`
	Amount              float64 `json:"amount"`
}
