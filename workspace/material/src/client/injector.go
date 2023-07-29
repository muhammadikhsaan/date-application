package client

import (
	"context"

	"pensiel.com/material/src/client/postgresql"
)

type Client struct {
	Dbi postgresql.Client
}

func NewClient(ctx context.Context) (*Client, map[string]bool) {
	status := map[string]bool{}

	postgres, err := postgresql.NewClient()
	status["postgresql"] = true

	if err != nil {
		status["postgresql"] = false
	}

	return &Client{
		Dbi: postgres,
	}, status
}
