package rest

import (
	"net/http"

	"github.com/larlandis/shorturl/internal/pkg/config"
)

func Middle(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := config.WithContext(r.Context())
		next.ServeHTTP(w, r.Clone(ctx))
	}
	return http.HandlerFunc(fn)
}
