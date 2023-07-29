package privilages

import (
	"time"

	"pensiel.com/material/src/contract"
)

type Entity struct {
	UserID      uint       `json:"user_id,omitempty" gorm:"not null;"`
	Feature     string     `json:"feature,omitempty" gorm:"not null;"`
	ExpiredDate *time.Time `json:"expired_date,omitempty"`
}

type EntityModel struct {
	// Basic Entity
	contract.MetaEntity

	// Self Entity
	Entity
}

func (EntityModel) TableName() string {
	return "user_privilages"
}
