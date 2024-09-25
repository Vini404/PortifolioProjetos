package main

import (
	"github.com/go-chi/chi"
	"sync"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	customerController := ServiceContainer().InjectPlayerController()

	r := chi.NewRouter()
	r.HandleFunc("/list", customerController.ListCustomers)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
