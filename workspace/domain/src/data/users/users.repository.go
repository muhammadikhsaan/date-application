package users

import (
	"context"
	"net/http"

	"pensiel.com/material/src/client/postgresql"
	"pensiel.com/material/src/pensiel"
)

type Repository interface {
	SelectUserIdBySecondaryId(ctx context.Context, model *EntityModel) *pensiel.Error
	FindAllUserAbleInteraction(ctx context.Context, userId uint, limit int, model *[]EntityModel) *pensiel.Error
	SelectByEmailOrUsername(ctx context.Context, model *EntityModel) *pensiel.Error
	SelectBySecondaryId(ctx context.Context, model *EntityModel) *pensiel.Error
	Insert(ctx context.Context, model *EntityModel) *pensiel.Error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repository) FindAllUserAbleInteraction(ctx context.Context, userId uint, limit int, model *[]EntityModel) *pensiel.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	err := dbi.
		Where("id NOT IN (select ui.target_id from user_interactions ui where ui.owner_id = ? and TO_CHAR(ui.updated_at , 'ddmmyyyy') = TO_CHAR(now(), 'ddmmyyyy'))", userId).
		Where("id != ?", userId).
		Order("random()").
		Limit(limit).
		Find(&model).
		Error

	if err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Origin:     err,
			Message:    "failed to find all user able",
		}
	}

	return nil
}

func (repository) SelectUserIdBySecondaryId(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	err := dbi.
		Where("users.secondary_id = ?", model.SecondaryId).
		First(&model).Error

	if err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Origin:     err,
			Message:    "failed to get user id using secondary id",
		}
	}

	return nil
}

func (repository) SelectBySecondaryId(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	err := dbi.
		Preload("Privilages").
		Where("users.secondary_id = ?", model.SecondaryId).
		First(&model).Error

	if err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Origin:     err,
			Message:    "failed to get user using secondary id",
		}
	}

	return nil
}

func (repository) SelectByEmailOrUsername(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	err := dbi.
		Where(
			dbi.Where("email = ?", model.Email).
				Or("username = ?", model.Username),
		).
		First(&model).Error

	if err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Origin:     err,
			Message:    "failed to get user using user credentioal",
		}
	}

	return nil
}

func (repository) Insert(ctx context.Context, model *EntityModel) *pensiel.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	if err := dbi.Create(model).Error; err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to insert users into database",
		}
	}

	return nil
}
