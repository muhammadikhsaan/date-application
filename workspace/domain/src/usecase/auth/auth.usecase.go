package auth

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
	"pensiel.com/domain/src/data/users"
	"pensiel.com/material/src/client"
	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/helper"
	"pensiel.com/material/src/modules"
	"pensiel.com/material/src/pensiel"
	"pensiel.com/material/src/static"
)

type Usecase interface {
	Login(ctx context.Context, data ParamAuthLogin) (*string, *pensiel.Error)
	Register(ctx context.Context, data ParamAuthRegister) *pensiel.Error
	Me(ctx context.Context, me ParamAuthMe) (*users.EntityModel, *pensiel.Error)
	CheckAlredyUserCredential(ctx context.Context, me ParamAuthCheckAlredyUserCredential) *pensiel.Error
}

type usecase struct {
	*client.Client
	*Repository
}

type Repository struct {
	User users.Repository
}

func NewService(c *client.Client, r *Repository) Usecase {
	return &usecase{
		Client:     c,
		Repository: r,
	}
}

func (u *usecase) Register(ctx context.Context, data ParamAuthRegister) *pensiel.Error {
	return u.Dbi.Trx(ctx, func(tx context.Context) *pensiel.Error {

		user := &users.EntityModel{
			Entity: users.Entity{
				Email:    data.Email,
				Username: data.Username,
			},
		}

		if err := u.User.SelectByEmailOrUsername(tx, user); err != nil {
			if !errors.Is(err.Origin, gorm.ErrRecordNotFound) {
				return err
			}
		}

		if user.SecondaryId != "" {
			return &pensiel.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "user already exist",
			}
		}

		user.Password = data.Password
		user.Username = data.Username

		if err := u.User.Insert(tx, user); err != nil {
			return err
		}

		return nil
	})
}

func (u *usecase) Login(ctx context.Context, data ParamAuthLogin) (*string, *pensiel.Error) {
	c := u.Dbi.Cnx(ctx)

	user := &users.EntityModel{
		Entity: users.Entity{
			Email:    data.Username,
			Username: data.Username,
		},
	}

	if err := u.User.SelectByEmailOrUsername(c, user); err != nil {
		return nil, err
	}

	_, err := helper.Compare(user.Password, data.Password)

	if err != nil {
		return nil, err
	}

	token, err := modules.GenerateJWT(static.JWT_LOGIN_SUBJECT, &contract.UserFormToken{
		SecondaryId: user.SecondaryId,
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *usecase) Me(ctx context.Context, me ParamAuthMe) (*users.EntityModel, *pensiel.Error) {
	cn := u.Dbi.Cnx(ctx)

	user := &users.EntityModel{
		MetaEntity: contract.MetaEntity{
			ShowableEntity: contract.ShowableEntity{
				SecondaryId: me.SecondaryId,
			},
		},
	}

	if err := u.User.SelectBySecondaryId(cn, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *usecase) CheckAlredyUserCredential(ctx context.Context, data ParamAuthCheckAlredyUserCredential) *pensiel.Error {
	c := u.Dbi.Cnx(ctx)

	user := &users.EntityModel{
		Entity: users.Entity{
			Email:    data.Email,
			Username: data.Username,
		},
	}

	if err := u.User.SelectByEmailOrUsername(c, user); err != nil {
		if !errors.Is(err.Origin, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if user.SecondaryId != "" {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "user already exist",
		}
	}

	return nil
}
