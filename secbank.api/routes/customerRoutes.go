package routes

import (
	"github.com/go-chi/chi"
	"secbank.api/auth"
	"secbank.api/controllers"
	"secbank.api/database"
	"secbank.api/infrastructures"
	"secbank.api/repositories"
	"secbank.api/services"
)

type CustomerRoutes struct {
}

func (c CustomerRoutes) AddToRouter(r *chi.Mux) {

	var customerController = GetCustomerController()

	r.Post("/login", customerController.Auth)
	r.Post("/customer", customerController.Create)

	r.With(auth.AuthMiddleware).Get("/customer", customerController.List)
	r.With(auth.AuthMiddleware).Get("/customer/{id}", customerController.Get)
	r.With(auth.AuthMiddleware).Put("/customer", customerController.Update)
	r.With(auth.AuthMiddleware).Delete("/customer/{id}", customerController.Delete)
}

func GetCustomerController() controllers.CustomerController {
	sqliteHandler := &infrastructures.SQLHandler{}
	sqliteHandler.Conn = database.NewConnection()

	customerRepository := &repositories.CustomerRepository{sqliteHandler}
	accountRepository := &repositories.AccountRepository{sqliteHandler}
	accountHolderRepository := &repositories.AccountHolderRepository{sqliteHandler}
	balanceRepository := &repositories.BalanceRepository{sqliteHandler}

	customerService := &services.CustomerService{ICustomerRepository: customerRepository, IAccountRepository: accountRepository, IAccountHolderRepository: accountHolderRepository, IBalanceRepository: balanceRepository}

	customerController := controllers.CustomerController{ICustomerService: customerService}

	return customerController
}
