package dto

type TransferRequest struct {
	IDCreditAccount         int     `json:"credit_account"`
	IDCustomerOriginAccount int     `json:"id_customer_origin_account"`
	Amount                  float64 `json:"amount"`
}
