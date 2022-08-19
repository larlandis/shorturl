package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/larlandis/shorturl/internal/pkg/config"
	"github.com/larlandis/shorturl/internal/pkg/metrics"
)

func Middle(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := config.WithContext(r.Context())
		host, _ := os.Hostname()

		// create segment
		s := metrics.Segment(ctx, "request-time",
			metrics.Tag("origin-host", host),
		)
		w.Header().Add("X-Origin-Host", host)

		next.ServeHTTP(w, r.Clone(ctx))

		route := fmt.Sprintf("[%s] %s", r.Method, chi.RouteContext(r.Context()).RoutePattern())
		metrics.Count(ctx, "request-count",
			metrics.Tag("route", route),
			metrics.Tag("origin-host", host),
		)
		s.End(metrics.Tag("route", route))

	}
	return http.HandlerFunc(fn)
}
