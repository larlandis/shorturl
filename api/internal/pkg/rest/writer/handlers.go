package writer

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"github.com/larlandis/shorturl/internal/pkg/config"
	"github.com/larlandis/shorturl/internal/pkg/errors"
	"github.com/larlandis/shorturl/internal/pkg/rest"
)

const hashLenKey = "hash.length"

type (
	RequestBody struct {
		Input string `json:"input"`
	}
	Response struct {
		Short string `json:"short"`
	}
)

func (s *Server) createHash() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// unmarshall body
		var data RequestBody
		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			rest.Error(w, errors.JSONDecodingError(err))
			return
		}
		if data.Input == "" {
			rest.Error(w, errors.InvalidInputError)
			return
		}

		// create hash
		ctx := req.Context()
		hash, err := s.service.CreateNewHash(
			ctx, data.Input, config.GetInt(ctx, hashLenKey),
		)
		if err != nil {
			rest.Error(w, err)
			return
		}
		render.JSON(w, req, Response{Short: hash})
	}
}
