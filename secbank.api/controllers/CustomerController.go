package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"secbank.api/interfaces"
)

type CustomerController struct {
	interfaces.ICustomerService
}

func (controller *CustomerController) ListCustomers(res http.ResponseWriter, req *http.Request) {

	customers, err := controller.List()
	if err != nil {
		fmt.Println(err.Error())
	}

	json.NewEncoder(res).Encode(customers)
}
