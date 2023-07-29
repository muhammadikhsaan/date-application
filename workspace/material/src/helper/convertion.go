package helper

import (
	"encoding/json"
	"net/http"

	"pensiel.com/material/src/pensiel"
)

func ToValue[T comparable](value *T) T {
	var v T

	if value != nil {
		return *value
	}

	return v
}

func ToPointer[T comparable](value T) *T {
	return &value
}

func ToMap(v any) (map[string]any, *pensiel.Error) {
	var val map[string]any

	j, err := json.Marshal(v)

	if err != nil {
		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "Failed to marshal json",
		}
	}

	if err := json.Unmarshal(j, &val); err != nil {
		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "Failed to unmarshal json",
		}
	}

	return val, nil
}

func ToStruct(m map[string]any, v any) *pensiel.Error {
	j, err := json.Marshal(v)

	if err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "Failed to marshal json",
		}
	}

	if err := json.Unmarshal(j, v); err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "Failed to unmarshal json",
		}
	}

	return nil
}
