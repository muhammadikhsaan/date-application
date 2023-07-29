package middleware

import (
	"net/http"

	c "github.com/go-chi/cors"
)

func Cors(next http.Handler) http.Handler {
	return c.Handler(c.Options{
		// AllowedOrigins:   []string{"https://foo.com"},
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{""},
		AllowCredentials: false,
		MaxAge:           300,
	})(next)
}
