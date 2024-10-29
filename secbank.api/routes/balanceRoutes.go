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

type BalanceRoutes struct {
}

func (c BalanceRoutes) AddToRouter(r *chi.Mux) {

	var balanceController = GetBalanceController()

	r.With(auth.AuthMiddleware).Get("/balance/{accountID}", balanceController.Get)
}

func GetBalanceController() controllers.BalanceController {
	sqliteHandler := &infrastructures.SQLHandler{}
	sqliteHandler.Conn = database.NewConnection()

	accountRepository := &repositories.BalanceRepository{sqliteHandler}

	accountService := &services.BalanceService{accountRepository}
	accountController := controllers.BalanceController{IBalanceService: accountService}

	return accountController
}
