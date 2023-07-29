package data

import (
	"pensiel.com/domain/src/data/interactions"
	"pensiel.com/domain/src/data/privilages"
	"pensiel.com/domain/src/data/users"
)

type Repository struct {
	User        users.Repository
	Interaction interactions.Repository
	Privilages  privilages.Repository
}

func NewRepository() *Repository {
	return &Repository{
		User:        users.NewRepository(),
		Interaction: interactions.NewRepository(),
		Privilages:  privilages.NewRepository(),
	}
}
