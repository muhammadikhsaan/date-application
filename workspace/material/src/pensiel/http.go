package pensiel

import (
	ctx "context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"pensiel.com/material/src/contract"
)

type HandlerFunc func(Context) *Error

type Context interface {
	User() *contract.UserFormToken
	Bind(v any) *Error
	BasicAuth() (string, string, bool)
	Param(k string) string
	Context() ctx.Context
	Query(k string) string
	Header() http.Header
	JSON(code int, i any) *Error
	FILE(code int, extention string, i io.Reader) *Error
	Validate(v any) *Error
	FormValue(key string) string
	FormFile(key string) (multipart.File, *multipart.FileHeader, *Error)
	GetCookie(name string) (*http.Cookie, *Error)
	SetCookie(cookie *http.Cookie) *context
}

type context struct {
	file     File
	request  *http.Request
	response http.ResponseWriter
	validate *validator.Validate
}

func New(w http.ResponseWriter, r *http.Request) Context {
	v := validator.New()
	f := NewFile()

	return &context{
		response: w,
		request:  r,
		validate: v,
		file:     f,
	}
}

func (c *context) User() *contract.UserFormToken {
	user := new(contract.UserFormToken)
	return c.Context().Value(reflect.TypeOf(user)).(*contract.UserFormToken)
}

func (c *context) GetCookie(name string) (*http.Cookie, *Error) {
	cookie, err := c.request.Cookie(name)

	if err != nil {
		return nil, &Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to get cookie form this request",
		}
	}

	return cookie, nil
}

func (c *context) FormFile(key string) (multipart.File, *multipart.FileHeader, *Error) {
	file, header, err := c.request.FormFile(key)

	if err != nil {
		return nil, nil, &Error{
			StatusCode: http.StatusBadRequest,
			Origin:     err,
			Message:    "file was not submitted in this request",
		}
	}

	return file, header, nil
}

func (c *context) FormValue(key string) string {
	return c.request.FormValue(key)
}

func (c *context) Context() ctx.Context {
	return c.request.Context()
}

func (c *context) Param(k string) string {
	return chi.URLParam(c.request, k)
}

func (c *context) Query(k string) string {
	return c.request.URL.Query().Get(k)
}

func (c *context) Header() http.Header {
	return c.request.Header
}

func (c *context) BasicAuth() (string, string, bool) {
	return c.request.BasicAuth()
}

func (c *context) Bind(v any) *Error {
	if err := json.NewDecoder(c.request.Body).Decode(v); err != nil {
		return &Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed binding this request body",
		}
	}

	return nil
}

func (c *context) Validate(v any) *Error {
	if err := c.validate.Struct(v); err != nil {
		return &Error{
			StatusCode: http.StatusBadRequest,
			Origin:     err,
			Message:    "the request is not in accordance with the contract",
		}
	}

	return nil
}

func (c *context) JSON(code int, i any) *Error {
	c.response.Header().Set("Content-Type", "application/json")
	c.response.WriteHeader(code)

	if err := json.NewEncoder(c.response).Encode(i); err != nil {
		return &Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to encode response",
		}
	}

	return nil
}

func (c *context) FILE(code int, extention string, file io.Reader) *Error {

	if err := c.file.SetHeader(c.response, extention); err != nil {
		return err
	}

	c.response.WriteHeader(code)

	if _, err := io.Copy(c.response, file); err != nil {
		return &Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to write file response",
		}
	}

	return nil
}

func (c *context) SetCookie(cookie *http.Cookie) *context {
	http.SetCookie(c.response, cookie)
	return c
}
