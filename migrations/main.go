package main

import (
	"context"
	"fmt"

	"pensiel.com/material/src/client/postgresql"

	"pensiel.com/material/src/client/postgresql/migrate"
)

func main() {
	ctx := context.Background()
	c, _ := postgresql.NewClient()

	dbi := c.Cnx(ctx).(*postgresql.Connection).Conn

	if err := migrate.Interaction(dbi).Executor(); err != nil {
		fmt.Println(err)
	}

	if err := migrate.Privilage(dbi).Executor(); err != nil {
		fmt.Println(err)
	}

	if err := migrate.User(dbi).Executor(); err != nil {
		fmt.Println(err)
	}

	if err := migrate.User(dbi).Seeder(); err != nil {
		fmt.Println(err)
	}
}
