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

type TransactionRoutes struct {
}

func (c TransactionRoutes) AddToRouter(r *chi.Mux) {

	var transactionController = GetTransactionController()

	r.With(auth.AuthMiddleware).Post("/transaction", transactionController.Transfer)
}

func GetTransactionController() controllers.TransactionController {
	sqliteHandler := &infrastructures.SQLHandler{}
	sqliteHandler.Conn = database.NewConnection()

	balanceRepository := &repositories.BalanceRepository{sqliteHandler}
	accountRepository := &repositories.AccountRepository{sqliteHandler}
	transactionRepository := &repositories.TransactionRepository{sqliteHandler}
	customerRepository := &repositories.CustomerRepository{sqliteHandler}

	transactionService := &services.TransactionService{accountRepository, balanceRepository, transactionRepository, customerRepository}
	transactionController := controllers.TransactionController{transactionService}

	return transactionController
}
