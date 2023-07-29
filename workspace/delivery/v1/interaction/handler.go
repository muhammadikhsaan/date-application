package interaction

import (
	"net/http"

	"pensiel.com/domain/src/usecase/interaction"
	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

func (h *handler) USERABLE(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	user := c.User()

	data, err := h.uc.Interaction.GetUserInteraction(ctx, interaction.ParamGetUserInteraction{
		UserID: user.SecondaryId,
	})

	if err != nil {
		return err
	}

	resp := []UserAbleDataResponse{}

	for _, v := range data {
		r := UserAbleDataResponse{}
		r.MapFromEntity(v)
		resp = append(resp, r)
	}

	return c.JSON(http.StatusOK, UserAbleResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "successfully get users",
		},
		Data: resp,
	})
}

func (h *handler) ACTION(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	user := c.User()

	req := ActionRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if !req.Action.IsValid() {
		return &pensiel.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid action type",
		}
	}

	err := h.uc.Interaction.SaveInteractionAction(ctx, interaction.ParamSaveInteraction{
		UserID:   user.SecondaryId,
		TargetID: req.Target,
		Action:   string(req.Action),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, ActionResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "successfully insert new interaction",
		},
	})
}
