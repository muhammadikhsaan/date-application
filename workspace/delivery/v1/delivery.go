package v1

import (
	"pensiel.com/delivery/v1/auth"
	"pensiel.com/delivery/v1/interaction"
	"pensiel.com/delivery/v1/purchase"
	"pensiel.com/domain/src/usecase"
	"pensiel.com/material/src/middleware"
	"pensiel.com/material/src/pensiel"
)

type Delivery interface {
	Router(r pensiel.Router)
}

type delivery struct {
	uc *usecase.Service
}

func NewDelivery(uc *usecase.Service) Delivery {
	return &delivery{
		uc: uc,
	}
}

func (c *delivery) Router(r pensiel.Router) {
	// GUEST
	r.Group(func(r pensiel.Router) {
		r.Use(middleware.Throttle)
		r.Route("/auth", auth.NewHandler(c.uc).Router)
	})

	// USERS
	r.Group(func(r pensiel.Router) {
		r.Use(middleware.Throttle)
		r.Use(middleware.Accessable)
		r.Route("/interaction", interaction.NewHandler(c.uc).Router)
		r.Route("/purchase", purchase.NewHandler(c.uc).Router)
	})
}
