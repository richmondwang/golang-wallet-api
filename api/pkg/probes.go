package pkg

import (
	"github.com/go-chi/chi/v5"
	// "github.com/heptiolabs/healthcheck"
)

func probeHandlers() *chi.Mux {
	r := chi.NewRouter()
	// health := healthcheck.NewHandler()
	// set probe handlers
	// health.AddReadinessCheck("namespace-modules-refreshed", healthcheck.Async(handler, 5*time.Second))
	// health.AddLivenessCheck("namespace-modules-refreshed", healthcheck.Async(handler, 5*time.Second))
	// r.Get("/live", health.LiveEndpoint)
	// r.Get("/ready", health.ReadyEndpoint)
	return r
}
