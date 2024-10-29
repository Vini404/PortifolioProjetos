package repositories

import (
	dto "secbank.api/dto/balance"
	"secbank.api/interfaces"
	"secbank.api/models"
)

type BalanceRepository struct {
	interfaces.IDbHandler
}

func (repository *BalanceRepository) R_GetByAccountID(accountID int) (*models.Balance, error) {
	balance := models.Balance{}
	err := repository.QueryWithParamSingleRow("SELECT * FROM Balance WHERE idaccount=$1", &balance, accountID)

	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (repository *BalanceRepository) R_Extract(accountID int) (*dto.ExtractResponse, error) {
	extract := dto.ExtractResponse{}

	sql := `
			select t.description as operation_name, t.amount,concat('Enviado para ',c.fullname,' - Conta',a.number,'-',a.digit) as transfer_type  from transactions t
			inner join account a on a.id = t.idcreditaccount
			inner join accountholder ah on ah.id = a.idaccountholder 
			inner join customer c on c.id  = ah.idcustomer 
			where t.iddebitaccount = $1
			union
			select t2.description as operation_name, t2.amount,concat('Recebido de ',c.fullname,' - Conta',a.number,'-',a.digit) as transfer_type from transactions t2 
			inner join account a on a.id = t2.iddebitaccount
			inner join accountholder ah on ah.id = a.idaccountholder 
			inner join customer c on c.id  = ah.idcustomer 
			where idcreditaccount = $1
`
	err := repository.QueryWithParamSingleRow(sql, &extract, accountID)

	if err != nil {
		return nil, err
	}
	return &extract, nil
}

func (repository *BalanceRepository) R_Create(balance models.Balance) (int, error) {
	id, err := repository.Insert(balance, "balance")
	if err != nil {
		return 0, err
	}

	return id, nil
}
