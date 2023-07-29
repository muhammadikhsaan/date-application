package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"pensiel.com/domain/src/data"
	"pensiel.com/domain/src/usecase"
	"pensiel.com/material/src/client"
	pm "pensiel.com/material/src/middleware"
	"pensiel.com/material/src/pensiel"
	"pensiel.com/material/src/secret"
	"pensiel.com/material/src/static"

	v1 "pensiel.com/delivery/v1"
)

var (
	MODE = static.MODE
	PORT = fmt.Sprintf(":%s", static.PORT)
)

func init() {
	secret.LoadSecretKeyJWT()
}

func main() {
	r := pensiel.NewRouter()
	ctx := context.Background()

	if MODE == "DEBUG" {
		r.Use(middleware.Logger)
	}

	r.Use(pm.Recovery)
	r.Use(middleware.CleanPath)
	r.Use(pm.Cors)

	c, status := client.NewClient(ctx)
	dt := data.NewRepository()
	uc := usecase.NewService(dt, c)

	r.Route("/api/v1", v1.NewDelivery(uc).Router)

	fmt.Printf("Client Status : %v \n", status)
	fmt.Printf("Server running on port%s \n", PORT)
	if err := http.ListenAndServe(PORT, r); err != nil {
		ctx.Done()
		fmt.Println(err.Error())
	}
}
