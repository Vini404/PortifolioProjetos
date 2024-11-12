package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/jinzhu/copier"
	"net/http"
	"secbank.api/auth"
	"secbank.api/dto/customer"
	"secbank.api/interfaces/service"
	"secbank.api/models"
	"strconv"
	"strings"
	"time"
)

type CustomerController struct {
	interfaces.ICustomerService
}

func (controller *CustomerController) Auth(res http.ResponseWriter, req *http.Request) {
	var authRequest dto.AuthRequest

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&authRequest)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	token, err := controller.S_Auth(authRequest)

	if err != nil {
		SetResponseError(res, err)
		return
	}

	SetResponseSuccessWithPayload(res, token)

}

func (controller *CustomerController) List(res http.ResponseWriter, req *http.Request) {

	customers, err := controller.S_List()

	if err != nil {
		SetResponseError(res, err)
		return
	}

	SetResponseSuccessWithPayload(res, customers)
}

func (controller *CustomerController) Create(res http.ResponseWriter, req *http.Request) {
	// Define um limite de memória para o arquivo de upload
	err := req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		SetResponseError(res, err)
		return
	}

	// Extrai o arquivo de imagem
	imageUser, _, err := req.FormFile("file")
	if err != nil {
		SetResponseError(res, err)
		return
	}
	defer imageUser.Close()

	// Extrai os demais campos do formulário
	var customerRequest dto.CreateCustomerRequest
	customerRequest.FullName = req.FormValue("FullName")
	customerRequest.Phone = req.FormValue("Phone")
	customerRequest.Email = req.FormValue("Email")
	customerRequest.Password = req.FormValue("Password")
	customerRequest.Document = req.FormValue("Document")

	// Converte o campo Birthday de string para time.Time usando RFC3339
	birthdayStr := req.FormValue("Birthday")
	birthday, err := time.Parse(time.RFC3339, birthdayStr) // Formato esperado: YYYY-MM-DDTHH:MM:SS.sssZ
	if err != nil {
		SetResponseError(res, err)
		return
	}
	customerRequest.Birthday = birthday

	// Converte o DTO para o modelo de dados
	var customer models.Customer
	err = copier.Copy(&customer, &customerRequest)
	if err != nil {
		SetResponseError(res, err)
		return
	}

	// Chama o serviço para salvar o cliente com a imagem
	errInsert := controller.S_Create(customer, imageUser)
	if errInsert != nil {
		SetResponseError(res, errInsert)
		return
	}

	SetResponseSuccess(res)
}

func (controller *CustomerController) Update(res http.ResponseWriter, req *http.Request) {
	var customer models.Customer
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)

	errDecode := decoder.Decode(&customer)

	if errDecode != nil {
		SetResponseError(res, errDecode)
		return
	}

	updateErr := controller.S_Update(customer)

	if updateErr != nil {
		SetResponseError(res, updateErr)
		return
	}

	SetResponseSuccess(res)
}

func (controller *CustomerController) Get(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		SetResponseError(res, err)
		return
	}
	customer, errGet := controller.S_Get(idParsed)

	if errGet != nil {
		SetResponseError(res, errGet)
		return
	}

	SetResponseSuccessWithPayload(res, customer)
}

func (controller *CustomerController) GetCustomerByToken(res http.ResponseWriter, req *http.Request) {

	authHeader := req.Header.Get("Authorization")

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

	customer, errGet := controller.S_Get(auth.GetCustomerIDByJwtToken(tokenString))

	if errGet != nil {
		SetResponseError(res, errGet)
		return
	}

	SetResponseSuccessWithPayload(res, customer)
}

func (controller *CustomerController) Delete(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	idParsed, err := strconv.Atoi(id)

	if err != nil {
		SetResponseError(res, err)
		return
	}
	errToDelete := controller.S_Delete(idParsed)

	if errToDelete != nil {
		SetResponseError(res, errToDelete)
		return
	}

	SetResponseSuccess(res)
}
