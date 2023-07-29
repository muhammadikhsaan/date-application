package auth

import (
	"pensiel.com/domain/src/usecase"
	"pensiel.com/material/src/middleware"
	"pensiel.com/material/src/pensiel"
)

type Handler interface {
	Router(r pensiel.Router)
}

type handler struct {
	uc *usecase.Service
}

func NewHandler(uc *usecase.Service) Handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) Router(r pensiel.Router) {
	h.routerGuest(r)

	r.Group(func(r pensiel.Router) {
		r.Use(middleware.Accessable)
		h.routerAccess(r)
	})
}

func (h *handler) routerGuest(r pensiel.Router) {
	r.Get("/", h.LOGIN)
	r.Post("/", h.REGISTER)

	r.Get("/check/email/{email}", h.CHECKEMAIL)
	r.Get("/check/username/{username}", h.CHECKUSERNAME)
}

func (h *handler) routerAccess(r pensiel.Router) {
	r.Get("/me", h.ME)
	r.Delete("/", h.LOGOUT)
}
