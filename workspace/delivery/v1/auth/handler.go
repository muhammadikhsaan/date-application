package auth

import (
	"net/http"
	"strings"
	"time"

	"pensiel.com/domain/src/data/privilages"
	"pensiel.com/domain/src/usecase/auth"
	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
	"pensiel.com/material/src/static"
)

func (h *handler) REGISTER(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()

	body := new(AuthRegisterRequest)

	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := c.Validate(body); err != nil {
		return err
	}

	if err := h.uc.Auth.Register(ctx, auth.ParamAuthRegister{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &AuthRegisterResponse{
		contract.ResponseMeta{
			Message: "user successfully registered",
		},
	})
}

func (h *handler) LOGIN(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()

	username, password, ok := c.BasicAuth()

	if !ok {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get user from the reqeust",
		}
	}

	if username == "" || password == "" {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid username or password",
		}
	}

	username = strings.ToLower(username)

	token, err := h.uc.Auth.Login(ctx, auth.ParamAuthLogin{
		Username: username,
		Password: password,
	})

	if err != nil {
		return err
	}

	return c.SetCookie(&http.Cookie{
		Name:     static.COOKIE_REQUEST_IDENTITY_KEY,
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
		Secure:   static.ENVIRONTMENT != "LOCAL",
		Expires:  time.Now().AddDate(0, static.JWT_EXPIRED_TOKEN, 0),
	}).JSON(http.StatusOK, &AuthLoginResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "user has successfully logged in",
		},
		Data: AuthLoginDataResponse{
			Token: *token,
		},
	})
}

func (h *handler) ME(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	user := c.User()

	data, err := h.uc.Auth.Me(ctx, auth.ParamAuthMe{
		SecondaryId: user.SecondaryId,
	})

	if err != nil {
		return err
	}

	data.Password = ""

	privilages := []privilages.Entity{}

	for _, v := range data.Privilages {
		p := v.Entity
		p.UserID = 0
		privilages = append(privilages, p)
	}

	return c.JSON(http.StatusOK, &AuthMeResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "success get user data",
		},
		Data: AuthMeDataResponse{
			Entity:         data.Entity,
			ShowableEntity: data.ShowableEntity,
			Privilages:     privilages,
		},
	})
}

func (h *handler) CHECKEMAIL(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	email := c.Param("email")

	if email == "" {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid email",
		}
	}

	if err := h.uc.Auth.CheckAlredyUserCredential(ctx, auth.ParamAuthCheckAlredyUserCredential{
		Email: email,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &AuthCheckUserResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "e-mail can be used",
		},
	})
}

func (h *handler) CHECKUSERNAME(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	username := c.Param("username")

	if username == "" {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid username",
		}
	}

	if err := h.uc.Auth.CheckAlredyUserCredential(ctx, auth.ParamAuthCheckAlredyUserCredential{
		Username: username,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &AuthCheckUserResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "username can be used",
		},
	})
}

func (h *handler) LOGOUT(c pensiel.Context) *pensiel.Error {
	return c.SetCookie(&http.Cookie{
		Name:     static.COOKIE_REQUEST_IDENTITY_KEY,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   static.ENVIRONTMENT != "LOCAL",
		Expires:  time.Now().AddDate(0, 0, -1),
	}).JSON(http.StatusOK, &AuthLogoutResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "the user has successfully logged out",
		},
	})
}
