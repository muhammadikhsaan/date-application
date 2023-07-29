package main_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pensiel.com/delivery/v1/auth"
	"pensiel.com/material/src/contract"
)

var token = &contract.UserFormToken{
	SecondaryId: "aynGK823985300PsnLH",
}

var user = auth.AuthRegisterRequest{
	Username: fmt.Sprintf("user-test-1-%d", time.Now().Unix()),
	Email:    fmt.Sprintf("user-test-1-%d@user.id", time.Now().Unix()),
	Password: "userpassword10",
}

type RequestRegister struct {
	request auth.AuthRegisterRequest
}

type CaseRegister struct {
	title    string
	input    RequestRegister
	expected any
}

var caseRegister = []CaseRegister{
	{
		title: "(postive) register user with normal case",
		input: RequestRegister{
			request: auth.AuthRegisterRequest{
				Username: user.Username,
				Email:    user.Email,
				Password: user.Password,
			},
		},
		expected: http.StatusCreated,
	},
	{
		title: "(negative) register user with duplicate email",
		input: RequestRegister{
			request: auth.AuthRegisterRequest{
				Username: fmt.Sprintf("user-test-2-%d", time.Now().Unix()),
				Email:    user.Email,
				Password: user.Password,
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) register user with duplicate username",
		input: RequestRegister{
			request: auth.AuthRegisterRequest{
				Username: user.Username,
				Email:    fmt.Sprintf("user-test-2-%d@user.id", time.Now().Unix()),
				Password: user.Password,
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) register user without email",
		input: RequestRegister{
			request: auth.AuthRegisterRequest{
				Username: fmt.Sprintf("user-test-3-%d", time.Now().Unix()),
				Password: user.Password,
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) register user without username",
		input: RequestRegister{
			request: auth.AuthRegisterRequest{
				Email:    fmt.Sprintf("user-test-4-%d@user.id", time.Now().Unix()),
				Password: user.Password,
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) register user without password",
		input: RequestRegister{
			request: auth.AuthRegisterRequest{
				Username: fmt.Sprintf("user-test-5-%d", time.Now().Unix()),
				Email:    fmt.Sprintf("user-test-5-%d@user.id", time.Now().Unix()),
			},
		},
		expected: http.StatusBadRequest,
	},
}

func TestRegister(t *testing.T) {
	for key, v := range caseRegister {
		body, err := json.Marshal(v.input.request)

		assert.Nil(t, err, fmt.Sprintf("%s (%d) : marshal body", v.title, key))

		req, err := http.NewRequest("POST", authRoute, bytes.NewReader(body))

		assert.Nil(t, err, fmt.Sprintf("%s (%d) : execute request", v.title, key))

		nr := httptest.NewRecorder()
		router.ServeHTTP(nr, req)

		assert.Equal(t, v.expected, nr.Code, fmt.Sprintf("%s (%d) : response", v.title, key))
	}
}

type RequestDataLogin struct {
	Username string
	Password string
}

type RequestLogin struct {
	request RequestDataLogin
}

type CaseLogin struct {
	title    string
	input    RequestLogin
	expected any
}

var casesLogin = []CaseLogin{
	{
		title: "(postive) login user with normal case",
		input: RequestLogin{
			request: RequestDataLogin{
				Username: "useradmin",
				Password: "useradmin",
			},
		},
		expected: http.StatusOK,
	},
	{
		title: "(negative) login user with wrong username",
		input: RequestLogin{
			request: RequestDataLogin{
				Username: "useradmin-wrong",
				Password: "useradmin",
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) login user with wrong password",
		input: RequestLogin{
			request: RequestDataLogin{
				Username: "useradmin",
				Password: "useradmin-wrong",
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) login user with no username",
		input: RequestLogin{
			request: RequestDataLogin{
				Password: "useradmin",
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) login user with no password",
		input: RequestLogin{
			request: RequestDataLogin{
				Username: "useradmin",
			},
		},
		expected: http.StatusBadRequest,
	},
}

func TestLogin(t *testing.T) {
	for key, v := range casesLogin {
		req, err := http.NewRequest("GET", authRoute, nil)
		assert.Nil(t, err, fmt.Sprintf("%s (%d) : execute request", v.title, key))

		cred := fmt.Sprintf("%s:%s", v.input.request.Username, v.input.request.Password)
		auth := base64.StdEncoding.EncodeToString([]byte(cred))
		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))

		nr := httptest.NewRecorder()
		router.ServeHTTP(nr, req)

		assert.Equal(t, v.expected, nr.Code, fmt.Sprintf("%s (%d) : response", v.title, key))
	}
}
