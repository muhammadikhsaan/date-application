package usecase

import (
	"pensiel.com/domain/src/data"
	"pensiel.com/domain/src/usecase/auth"
	"pensiel.com/domain/src/usecase/interaction"
	"pensiel.com/domain/src/usecase/purchase"
	"pensiel.com/material/src/client"
)

type Service struct {
	Auth        auth.Usecase
	Purchase    purchase.Usecase
	Interaction interaction.Usecase
}

func NewService(r *data.Repository, c *client.Client) *Service {
	return &Service{
		Auth: auth.NewService(c, &auth.Repository{
			User: r.User,
		}),
		Purchase: purchase.NewService(c, &purchase.Repository{
			User:       r.User,
			Privilages: r.Privilages,
		}),
		Interaction: interaction.NewService(c, &interaction.Repository{
			User:        r.User,
			Interaction: r.Interaction,
		}),
	}
}
