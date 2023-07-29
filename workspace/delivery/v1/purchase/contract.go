package purchase

import "pensiel.com/material/src/contract"

type (
	Feature string

	PurchasePrivilagesRequest struct {
		Feature Feature `json:"feature" validate:"required"`
	}

	PurchasePrivilagesResponse struct {
		contract.ResponseMeta
	}
)

func (f Feature) IsValid() bool {
	return f == "quota" || f == "verify"
}
