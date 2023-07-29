package purchase

import (
	"context"
	"net/http"

	"pensiel.com/domain/src/data/privilages"
	"pensiel.com/domain/src/data/users"
	"pensiel.com/material/src/client"
	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

type Usecase interface {
	PurchasePrivilages(ctx context.Context, params *ParamPurchasePrivilages) *pensiel.Error
}

type usecase struct {
	*client.Client
	*Repository
}

type Repository struct {
	User       users.Repository
	Privilages privilages.Repository
}

func NewService(c *client.Client, r *Repository) Usecase {
	return &usecase{
		Client:     c,
		Repository: r,
	}
}

func (uc *usecase) PurchasePrivilages(ctx context.Context, params *ParamPurchasePrivilages) *pensiel.Error {
	return uc.Dbi.Trx(ctx, func(tx context.Context) *pensiel.Error {
		user := &users.EntityModel{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: params.UserID,
				},
			},
		}

		if err := uc.User.SelectUserIdBySecondaryId(tx, user); err != nil {
			return err
		}

		if !user.IsExist() {
			return &pensiel.Error{
				Message:    "users not found",
				StatusCode: http.StatusBadRequest,
			}
		}

		if err := uc.Privilages.Insert(tx, &privilages.EntityModel{
			Entity: privilages.Entity{
				UserID:      user.ID,
				Feature:     params.Feature,
				ExpiredDate: nil,
			},
		}); err != nil {
			return err
		}

		return nil
	})
}
