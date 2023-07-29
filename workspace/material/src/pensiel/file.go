package pensiel

import (
	"net/http"
)

type File interface {
	SetHeader(w http.ResponseWriter, extention string) *Error
}

type file struct {
	extentions map[string]string
}

func NewFile() File {
	extentions := map[string]string{
		"pdf":  "application/pdf",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
	}

	return &file{
		extentions: extentions,
	}
}

func (f *file) SetHeader(w http.ResponseWriter, extention string) *Error {
	val, ok := f.extentions[extention]

	if !ok {
		return &Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "File extention is not supported",
		}
	}

	w.Header().Set("Content-Type", val)
	return nil
}
