package reader

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/larlandis/shorturl/internal/pkg/errors"
	"github.com/larlandis/shorturl/internal/pkg/rest"
)

type Response struct {
	Data string `json:"data"`
}

func (s *Server) getHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// unmarshall body
		hash := chi.URLParam(r, "hash")
		if hash == "" {
			rest.Error(w, errors.InvalidInputError)
		}

		// create hash
		ctx := r.Context()
		data, err := s.service.GetFromHash(ctx, hash)
		if err != nil {
			rest.Error(w, err)
			return
		}
		render.JSON(w, r, Response{Data: data})
	}
}
