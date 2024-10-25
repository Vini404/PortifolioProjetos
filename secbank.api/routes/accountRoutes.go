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

type AccountRoutes struct {
}

func (c AccountRoutes) AddToRouter(r *chi.Mux) {

	var accountController = GetAccountController()

	r.With(auth.AuthMiddleware).Get("/account", accountController.List)
	r.With(auth.AuthMiddleware).Get("/account/{id}", accountController.Get)
	r.With(auth.AuthMiddleware).Post("/account", accountController.Create)
	r.With(auth.AuthMiddleware).Put("/account", accountController.Update)
	r.With(auth.AuthMiddleware).Delete("/account/{id}", accountController.Delete)
}

func GetAccountController() controllers.AccountController {
	sqliteHandler := &infrastructures.SQLHandler{}
	sqliteHandler.Conn = database.NewConnection()

	accountRepository := &repositories.AccountRepository{sqliteHandler}

	accountService := &services.AccountService{accountRepository}
	accountController := controllers.AccountController{IAccountService: accountService}

	return accountController
}
