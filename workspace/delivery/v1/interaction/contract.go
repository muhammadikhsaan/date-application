package interaction

import (
	"pensiel.com/domain/src/data/users"
	"pensiel.com/material/src/contract"
)

type (
	UserAbleDataResponse struct {
		SecondaryId string `json:"secondary_id"`
		Email       string `json:"email"`
	}

	UserAbleResponse struct {
		contract.ResponseMeta
		Data []UserAbleDataResponse `json:"data"`
	}
)

func (r *UserAbleDataResponse) MapFromEntity(entity users.EntityModel) {
	r.SecondaryId = entity.SecondaryId
	r.Email = entity.Email
}

type (
	Action string

	ActionRequest struct {
		Target string `json:"target" validate:"required"`
		Action Action `json:"action" validate:"required"`
	}

	ActionResponse struct {
		contract.ResponseMeta
	}
)

func (f Action) IsValid() bool {
	return f == "pass" || f == "like"
}
