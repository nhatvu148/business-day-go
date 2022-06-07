package middlewares

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func RequestPathLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("page", r.URL.Path).Msg("Endpoint Hit")
		next.ServeHTTP(w, r)
	})
}
