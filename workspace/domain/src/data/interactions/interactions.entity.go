package interactions

import (
	"pensiel.com/material/src/contract"
)

type Entity struct {
	OwnerID  uint   `json:"owner_id"`
	TargetID uint   `json:"target_id"`
	Action   string `json:"action"`
}

type EntityModel struct {
	// Basic Entity
	contract.MetaEntity

	// Self Entity
	Entity
}

func (EntityModel) TableName() string {
	return "user_interactions"
}
