package main

import (
	"secbank.api/controllers"
	"secbank.api/database"
	"secbank.api/infrastructures"
	"secbank.api/repositories"
	"secbank.api/services"
	"sync"
)

type IServiceContainer interface {
	InjectPlayerController() controllers.CustomerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.CustomerController {
	sqliteHandler := &infrastructures.SQLiteHandler{}
	sqliteHandler.Conn = database.NewConnection()

	customerRepository := &repositories.CustomerRepository{sqliteHandler}

	customerService := &services.CustomerService{customerRepository}
	customerController := controllers.CustomerController{customerService}

	return customerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
