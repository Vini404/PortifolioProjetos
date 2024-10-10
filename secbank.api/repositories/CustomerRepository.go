package repositories

import (
	"secbank.api/interfaces"
	"secbank.api/models"
	"secbank.api/utils"
)

type CustomerRepository struct {
	interfaces.IDbHandler
}

func (repository *CustomerRepository) R_List() (*[]models.Customer, error) {
	var customers []models.Customer
	err := repository.Query("SELECT * FROM Customer", &customers)

	if err != nil {
		return nil, err
	}

	// Loop through rows, using Scan to assign column data to struct fields.

	return &customers, nil
}

func (repository *CustomerRepository) R_Create(customer *models.Customer) error {
	err := repository.Insert(customer, "customer")

	if err != nil {
		return err
	}
	return nil
}

func (repository *CustomerRepository) R_Update(customer models.Customer) error {
	customerToUpdate := utils.StructToMap(customer)
	err := repository.Update(customer.ID, "customer", customerToUpdate)

	if err != nil {
		return err
	}
	return nil
}

func (repository *CustomerRepository) R_Delete(id int) error {

	err := repository.Delete(id, "customer")

	if err != nil {
		return err
	}
	return nil
}

func (repository *CustomerRepository) R_Get(id int) (*models.Customer, error) {
	customer := models.Customer{}
	err := repository.Get(id, "customer", &customer)

	if err != nil {
		return nil, err
	}
	return &customer, nil
}
