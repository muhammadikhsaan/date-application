package middleware

import (
	"net/http"

	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

func Recovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pensiel.New(w, r)

		defer func() {
			if r := recover(); r != nil {
				response := &contract.ResponseError{
					Message: "we have problem for this request",
				}

				if err, ok := r.(*pensiel.Error); ok {
					response.Message = err.Message

					if err.Origin != nil {
						response.Origin = err.Origin.Error()
					}
				}

				if err, ok := r.(error); ok {
					response.Origin = err.Error()
				}

				if err, ok := r.(string); ok {
					response.Origin = err
				}

				ctx.JSON(http.StatusInternalServerError, response)
				return
			}
		}()

		h.ServeHTTP(w, r)
	})
}
