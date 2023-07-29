package main_test

import (
	"context"
	"testing"

	"pensiel.com/delivery/v1/auth"
	"pensiel.com/delivery/v1/interaction"
	"pensiel.com/delivery/v1/purchase"
	"pensiel.com/domain/src/data"
	"pensiel.com/domain/src/usecase"
	"pensiel.com/material/src/client"
	"pensiel.com/material/src/pensiel"
	"pensiel.com/material/src/secret"
)

const (
	authRoute        = "/testing/auth"
	purchaseRoute    = "/testing/purchase"
	interactionRoute = "/testing/interaction"
)

var (
	router *pensiel.Mux
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	secret.LoadSecretKeyJWT()

	router = pensiel.NewRouter()

	c, _ := client.NewClient(ctx)

	dt := data.NewRepository()
	uc := usecase.NewService(dt, c)

	router.Route(authRoute, auth.NewHandler(uc).Router)
	router.Route(purchaseRoute, purchase.NewHandler(uc).Router)
	router.Route(interactionRoute, interaction.NewHandler(uc).Router)
	m.Run()
}
