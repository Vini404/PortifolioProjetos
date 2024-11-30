package repositories

import (
	"secbank.api/interfaces"
	"secbank.api/models"
	"secbank.api/utils"
	"time"
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

	return &customers, nil
}

func (repository *CustomerRepository) R_Create(customer models.Customer) (int, error) {
	customer.CreatedTimeStamp = time.Now()

	err := customer.Validate()

	if err != nil {
		return 0, err
	}

	id, err := repository.Insert(customer, "customer")
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repository *CustomerRepository) R_Update(customer models.Customer) error {
	customer.UpdatedTimeStamp = time.Now()
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

func (repository *CustomerRepository) R_Get_By_Email(email string) (*models.Customer, error) {
	customer := models.Customer{}

	err := repository.QueryWithParamSingleRow("SELECT * FROM Customer WHERE email=$1", &customer, email)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}
