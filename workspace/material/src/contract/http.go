package contract

type (
	ResponseMeta struct {
		Message    string              `json:"message"`
		Pagination *ResponsePagination `json:"pagination,omitempty"`
	}

	ResponsePagination struct {
		Current int64 `json:"current"`
		Limit   int64 `json:"limit"`
		Total   int64 `json:"total"`
	}

	ResponseError struct {
		Message string `json:"message,omitempty"`
		Origin  string `json:"origin,omitempty"`
	}
)
