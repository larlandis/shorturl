package rest

import (
	"errors"
	"net/http"

	apiErrors "github.com/larlandis/shorturl/internal/pkg/errors"
)

func Error(w http.ResponseWriter, err error) {
	var e apiErrors.ApiError
	if ok := errors.As(err, &e); !ok {
		e = apiErrors.New("internal server error", "unhandled error", http.StatusInternalServerError, err)
	}
	http.Error(w, e.Error(), e.Status())
}
