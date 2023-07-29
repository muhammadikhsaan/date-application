package main_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"pensiel.com/delivery/v1/interaction"
)

type CaseGetUserAble struct {
	title    string
	expected any
}

var getUserAbleCase = []CaseGetUserAble{
	{
		title:    "(positive) get user able",
		expected: http.StatusOK,
	},
}

func TestGetUserAble(t *testing.T) {
	for key, v := range getUserAbleCase {
		req, err := http.NewRequest("GET", interactionRoute, nil)

		assert.Nil(t, err, fmt.Sprintf("%s (%d) : execute request", v.title, key))

		var ctx = context.WithValue(req.Context(), reflect.TypeOf(token), token)

		nr := httptest.NewRecorder()
		router.ServeHTTP(nr, req.WithContext(ctx))

		assert.Equal(t, v.expected, nr.Code, fmt.Sprintf("%s (%d) : response", v.title, key))
	}
}

type RequestInteractioAction struct {
	request interaction.ActionRequest
}

type CaseInteractioAction struct {
	title    string
	input    RequestInteractioAction
	expected any
}

var interationActionCase = []CaseInteractioAction{
	{
		title: "(positive) like action",
		input: RequestInteractioAction{
			request: interaction.ActionRequest{
				Target: "aynGK823985300PsnLZ",
				Action: "like",
			},
		},
		expected: http.StatusCreated,
	},
	{
		title: "(positive) pass action",
		input: RequestInteractioAction{
			request: interaction.ActionRequest{
				Target: "aynGK823985300PsnLp",
				Action: "pass",
			},
		},
		expected: http.StatusCreated,
	},
	{
		title: "(negative) undefined action",
		input: RequestInteractioAction{
			request: interaction.ActionRequest{
				Target: "aynGK823985300PsnLp",
				Action: "undefined",
			},
		},
		expected: http.StatusBadRequest,
	},
	{
		title: "(negative) no target",
		input: RequestInteractioAction{
			request: interaction.ActionRequest{
				Target: "nousertarget",
				Action: "pass",
			},
		},
		expected: http.StatusBadRequest,
	},
}

func TestInteractionAction(t *testing.T) {
	for key, v := range interationActionCase {
		body, err := json.Marshal(v.input.request)
		assert.Nil(t, err, fmt.Sprintf("%s (%d) : marshal body", v.title, key))

		req, err := http.NewRequest("POST", interactionRoute, bytes.NewReader(body))

		assert.Nil(t, err, fmt.Sprintf("%s (%d) : execute request", v.title, key))

		ctx := context.WithValue(req.Context(), reflect.TypeOf(token), token)

		nr := httptest.NewRecorder()
		router.ServeHTTP(nr, req.WithContext(ctx))

		assert.Equal(t, v.expected, nr.Code, fmt.Sprintf("%s (%d) : response", v.title, key))
	}
}
