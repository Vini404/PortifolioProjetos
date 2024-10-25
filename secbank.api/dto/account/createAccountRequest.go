package dto

type CreateAccountRequest struct {
	IDAccountHolder int    `db:"id_account_holder"`
	Number          string `db:"number"`
	Digit           string `db:"digit"`
}
