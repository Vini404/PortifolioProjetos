package dto

type TransferRequest struct {
	DigitCreditAccount      string  `json:"digit_credit_account"`
	NumberCreditAccount     string  `json:"number_credit_account"`
	IDCustomerOriginAccount int     `json:"id_customer_origin_account"`
	Amount                  float64 `json:"amount"`
}
