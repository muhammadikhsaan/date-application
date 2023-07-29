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
	"pensiel.com/delivery/v1/purchase"
)

type RequestPurchase struct {
	request purchase.PurchasePrivilagesRequest
}

type CasePurchase struct {
	title    string
	input    RequestPurchase
	expected any
}

var purchaseCase = []CasePurchase{
	{
		title: "(positive) purchase verify privilages",
		input: RequestPurchase{
			request: purchase.PurchasePrivilagesRequest{
				Feature: "verify",
			},
		},
		expected: http.StatusCreated,
	},
	{
		title: "(positive) purchase quota privilages",
		input: RequestPurchase{
			request: purchase.PurchasePrivilagesRequest{
				Feature: "quota",
			},
		},
		expected: http.StatusCreated,
	},
	{
		title: "(negative) purchase undefined privilages",
		input: RequestPurchase{
			request: purchase.PurchasePrivilagesRequest{
				Feature: "undefined",
			},
		},
		expected: http.StatusBadRequest,
	},
}

func TestPurchase(t *testing.T) {
	for key, v := range purchaseCase {
		body, err := json.Marshal(v.input.request)
		assert.Nil(t, err, fmt.Sprintf("%s (%d) : marshal body", v.title, key))

		req, err := http.NewRequest("POST", purchaseRoute, bytes.NewReader(body))

		assert.Nil(t, err, fmt.Sprintf("%s (%d) : execute request", v.title, key))

		ctx := context.WithValue(req.Context(), reflect.TypeOf(token), token)

		nr := httptest.NewRecorder()
		router.ServeHTTP(nr, req.WithContext(ctx))

		assert.Equal(t, v.expected, nr.Code, fmt.Sprintf("%s (%d) : response", v.title, key))
	}
}
