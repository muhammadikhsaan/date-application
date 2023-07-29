package middleware

import (
	"context"
	"net/http"
	"reflect"

	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/modules"
	"pensiel.com/material/src/pensiel"
	"pensiel.com/material/src/static"
)

func Accessable(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := pensiel.New(w, r)

		cookie, err := c.GetCookie(static.COOKIE_REQUEST_IDENTITY_KEY)

		if err != nil {
			c.JSON(err.StatusCode, &contract.ResponseError{
				Message: err.Message,
				Origin:  err.Origin.Error(),
			})
			return
		}

		user, errs := modules.VerifyJWT(static.JWT_LOGIN_SUBJECT, cookie.Value)

		if errs != nil {
			resp := &contract.ResponseError{
				Message: errs.Message,
			}

			if errs.Origin != nil {
				resp.Origin = errs.Origin.Error()
			}

			c.JSON(http.StatusForbidden, resp)
			return
		}

		ctx := context.WithValue(c.Context(), reflect.TypeOf(user), user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
