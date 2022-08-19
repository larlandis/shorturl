package writer

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/larlandis/shorturl/internal/pkg/rest"
)

type (
	Server struct {
		chi.Router
		service service
	}
	service interface {
		CreateNewHash(ctx context.Context, input string, len uint) (string, error)
	}
)

func New(s service) *Server {
	r := chi.NewRouter()
	srv := &Server{
		Router:  r,
		service: s,
	}
	r.Use(middleware.Logger)
	r.Use(rest.Middle)
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/shorturl", srv.createHash())
	})
	return srv
}
