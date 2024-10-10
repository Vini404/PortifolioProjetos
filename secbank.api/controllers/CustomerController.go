package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"secbank.api/interfaces"
	"secbank.api/models"
	"strconv"
)

type CustomerController struct {
	interfaces.ICustomerService
}

func (controller *CustomerController) List(res http.ResponseWriter, req *http.Request) {

	customers, err := controller.S_List()
	if err != nil {
		fmt.Println(err.Error())
	}

	json.NewEncoder(res).Encode(customers)
}

func (controller *CustomerController) Create(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&customer)
	if errDecode != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	errInsert := controller.S_Create(customer)

	if errInsert != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
	}
}

func (controller *CustomerController) Update(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&customer)
	if errDecode != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	updateErr := controller.S_Update(customer)

	if updateErr != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
	}
}

func (controller *CustomerController) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	customer, err := controller.S_Get(idParsed)

	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
	}

	json.NewEncoder(res).Encode(customer)
}

func (controller *CustomerController) Delete(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	errToDelete := controller.S_Delete(idParsed)

	if errToDelete != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
	}
}
