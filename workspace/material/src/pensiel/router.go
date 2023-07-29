package pensiel

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"pensiel.com/material/src/contract"
)

type Router interface {
	Use(middlewares ...func(http.Handler) http.Handler)
	Route(pattern string, fn func(r Router))
	Group(fn func(r Router))
	Get(pattern string, h HandlerFunc)
	Post(pattern string, h HandlerFunc)
	Put(pattern string, h HandlerFunc)
	Patch(pattern string, h HandlerFunc)
	Delete(pattern string, h HandlerFunc)
}

type router struct {
	chi.Router
}

func (r *router) Use(middlewares ...func(http.Handler) http.Handler) {
	r.Router.Use(middlewares...)
}

func (r *router) Route(pattern string, fn func(r Router)) {
	r.Router.Route(pattern, func(r chi.Router) {
		fn(&router{r})
	})
}

func (r *router) Group(fn func(r Router)) {
	r.Router.Group(func(r chi.Router) {
		fn(&router{r})
	})
}

func (r *router) Get(pattern string, h HandlerFunc) {
	r.Router.Get(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := New(w, r)

		if err := h(ctx); err != nil {
			resp := &contract.ResponseError{
				Message: err.Message,
			}

			if err.Origin != nil {
				resp.Origin = err.Origin.Error()
			}

			ctx.JSON(err.StatusCode, resp)
			return
		}
	})
}

func (r *router) Put(pattern string, h HandlerFunc) {
	r.Router.Put(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := New(w, r)

		if err := h(ctx); err != nil {
			resp := &contract.ResponseError{
				Message: err.Message,
			}

			if err.Origin != nil {
				resp.Origin = err.Origin.Error()
			}

			ctx.JSON(err.StatusCode, resp)
			return
		}
	})
}

func (r *router) Patch(pattern string, h HandlerFunc) {
	r.Router.Patch(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := New(w, r)

		if err := h(ctx); err != nil {
			resp := &contract.ResponseError{
				Message: err.Message,
			}

			if err.Origin != nil {
				resp.Origin = err.Origin.Error()
			}

			ctx.JSON(err.StatusCode, resp)
			return
		}
	})
}

func (r *router) Delete(pattern string, h HandlerFunc) {
	r.Router.Delete(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := New(w, r)

		if err := h(ctx); err != nil {
			resp := &contract.ResponseError{
				Message: err.Message,
			}

			if err.Origin != nil {
				resp.Origin = err.Origin.Error()
			}

			ctx.JSON(err.StatusCode, resp)
			return
		}
	})
}

func (r *router) Post(pattern string, h HandlerFunc) {
	r.Router.Post(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := New(w, r)

		if err := h(ctx); err != nil {
			resp := &contract.ResponseError{
				Message: err.Message,
			}

			if err.Origin != nil {
				resp.Origin = err.Origin.Error()
			}

			ctx.JSON(err.StatusCode, resp)
			return
		}
	})
}
