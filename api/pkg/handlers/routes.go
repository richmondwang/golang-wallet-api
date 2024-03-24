package handlers

import (
	"github.com/go-chi/chi/v5"
)

// Routes setup routse for accounts handler
func (handlers *accounts) Routes(r chi.Router) chi.Router {
	r.Route("/accounts", func(r chi.Router) {
		r.Route("/{accountID}", func(r chi.Router) {
			r.Use(handlers.GetAccountMiddleware)

			r.Get("/", handlers.GetAccount)
			r.Post("/transfer", handlers.TransferMoney)
			r.Get("/transactions", handlers.Transactions)
		})

		r.Post("/", handlers.CreateAccount)
	})
	return r
}
