package purchase

import (
	"net/http"

	"pensiel.com/domain/src/usecase/purchase"
	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

func (h *handler) PURCHASE(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	user := c.User()

	req := PurchasePrivilagesRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if !req.Feature.IsValid() {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid feature type",
		}
	}

	if err := h.uc.Purchase.PurchasePrivilages(ctx, &purchase.ParamPurchasePrivilages{
		UserID:  user.SecondaryId,
		Feature: string(req.Feature),
	}); err != nil {
		return &pensiel.Error{}
	}

	return c.JSON(http.StatusCreated, PurchasePrivilagesResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "purchase success",
		},
	})
}
