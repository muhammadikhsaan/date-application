package privilages

import (
	"context"

	"pensiel.com/material/src/client/postgresql"
	"pensiel.com/material/src/pensiel"
)

type Repository interface {
	Insert(ctx context.Context, model *EntityModel) *pensiel.Error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	if err := dbi.Create(model).Error; err != nil {
		return &pensiel.Error{}
	}

	return nil
}
