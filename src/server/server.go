package server

import (
	"net/http"
	"time"
	"valsea/src/config"
	"valsea/src/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Server struct {
	router *chi.Mux
}

func NewServer(config *config.Config, accountHandler *handler.Account, transferHandler *handler.Transfer) *Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Duration(config.TimeOutInSeconds) * time.Second))

	AddRoutes(r, config.BaseUri, accountHandler, transferHandler)

	return &Server{
		router: r,
	}
}

func AddRoutes(router *chi.Mux, baseUri string, accountHandler *handler.Account, transferHandler *handler.Transfer) {
	router.Route(baseUri, func(r chi.Router) {
		r.Mount("/accounts", accountRouter(accountHandler))
		r.Mount("/transfer", transferRouter(transferHandler))

	})
}

func accountRouter(handler *handler.Account) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.CreateAccount)
	r.Get("/{id}", handler.GetAccountById)
	r.Get("/", handler.ListAccounts)
	r.Post("/{id}/transactions", handler.CreateTransaction)
	r.Get("/{id}/transactions", handler.GetTransactionsByAccountId)
	return r
}

func transferRouter(handler *handler.Transfer) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Transfer)
	return r
}

func (s *Server) ListenAndServe(address string) {
	zap.S().Infof("Server up and runing in port %v", address)
	if err := http.ListenAndServe(address, s.router); err != nil {
		zap.S().Errorf("Error starting server in addres: %v, %v", address, zap.Error(err))
	}
}
