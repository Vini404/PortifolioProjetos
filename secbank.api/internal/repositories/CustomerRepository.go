package repositories

import (
	"secbank.api/internal/interfaces"
	"secbank.api/internal/models"
)

type CustomerRepositoryWithCircuitBreaker struct {
	CustomerRepository interfaces.ICustomerRepository
}

type CustomerRepository struct {
	interfaces.IDbHandler
}

func (repository *CustomerRepository) List() (models.Customer, error) {

	row, err := repository.Query("SELECT * FROM Customer")
	if err != nil {
		return models.Customer{}, err
	}

	var player models.Customer

	row.Next()
	//row.Scan(&player.Id, &player.Name, &player.Score)

	return player, nil
}
