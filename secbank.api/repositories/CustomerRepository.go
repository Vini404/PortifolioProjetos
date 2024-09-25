package repositories

import (
	"secbank.api/interfaces"
	"secbank.api/models"
)

type CustomerRepository struct {
	interfaces.IDbHandler
}

func (repository *CustomerRepository) ListAllCustomer() (*[]models.Customer, error) {

	row, err := repository.Query("SELECT * FROM Customer")
	if err != nil {
		return nil, err
	}

	var customers []models.Customer

	// Loop through rows, using Scan to assign column data to struct fields.
	for row.Next() {
		var customer models.Customer
		if err := row.Scan(&customer.ID, &customer.FullName, &customer.Phone, &customer.Email, &customer.Birthday, &customer.CreatedTimeStamp, &customer.UpdatedTimeStamp); err != nil {
			return &customers, err
		}
		customers = append(customers, customer)
	}

	return &customers, nil
}
