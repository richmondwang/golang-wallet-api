package pkg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/richmondwang/golang-wallet-api/docs"
	"github.com/richmondwang/golang-wallet-api/ent"
	"github.com/richmondwang/golang-wallet-api/pkg/handlers"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// Server the API server
type Server struct {
	Port string
	DB   *ent.Client
}

// @title    		github.com/richmondwang/golang-wallet-api
// @version  		1
// @description     A simple wallet api.

// @contact.email   richrsw35@gmail.com
// @contact.name   	Richmond Wang

// @host      localhost:3000
// @BasePath  /
func (s *Server) Serve(ctx context.Context) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount("/healthz", probeHandlers())

	accountsHandler := handlers.NewAccountsHandler(s.DB)
	r.Mount("/", accountsHandler.Routes(r))

	r.Get("/swagger*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Port),
		Handler: r,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	fmt.Println("starting server at :8080")
	log.Fatal(server.ListenAndServe())
}
