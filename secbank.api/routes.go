package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"secbank.api/routes"
	"sync"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(corsMiddleware)
	r.Get("/healthcheck", HealthCheckHandler)

	routes.CustomerRoutes{}.AddToRouter(r)
	routes.AccountRoutes{}.AddToRouter(r)
	routes.BalanceRoutes{}.AddToRouter(r)
	routes.TransactionRoutes{}.AddToRouter(r)

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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
