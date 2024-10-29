package main

import (
	"github.com/go-chi/chi"
	"secbank.api/routes"
	"sync"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	r := chi.NewRouter()
	routes.CustomerRoutes{}.AddToRouter(r)
	routes.AccountRoutes{}.AddToRouter(r)
	routes.BalanceRoutes{}.AddToRouter(r)

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
