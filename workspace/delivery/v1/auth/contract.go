package auth

import (
	"pensiel.com/domain/src/data/privilages"
	"pensiel.com/domain/src/data/users"
	"pensiel.com/material/src/contract"
)

type (
	AuthRegisterRequest struct {
		Email    string `json:"email" validate:"required,email,max=124"`
		Username string `json:"username" validate:"required,min=6,max=124"`
		Password string `json:"password" validate:"required,min=8"`
	}

	AuthRegisterResponse struct {
		contract.ResponseMeta
	}
)

type (
	AuthLoginDataResponse struct {
		Token string `json:"token"`
	}

	AuthLoginResponse struct {
		contract.ResponseMeta
		Data AuthLoginDataResponse `json:"data"`
	}
)

type (
	AuthMeDataResponse struct {
		users.Entity
		contract.ShowableEntity
		Privilages []privilages.Entity
	}

	AuthMeResponse struct {
		contract.ResponseMeta
		Data AuthMeDataResponse `json:"data"`
	}
)

type (
	AuthCheckUserResponse struct {
		contract.ResponseMeta
	}
)

type (
	AuthLogoutResponse struct {
		contract.ResponseMeta
	}
)
