package contract

import (
	"time"

	"gorm.io/gorm"
	"pensiel.com/material/src/generator"
)

type SecureEntity struct {
	ID uint `json:"id,omitempty" gorm:"primaryKey;autoIncrement;index:idx_search_id,priority:1;"`

	DeletedAt gorm.DeletedAt `json:"-" gorm:"index:idx_search_id,priority:2;index:idx_search_secondary_id,priority:2;"`
}

type ShowableEntity struct {
	SecondaryId string `json:"secondary_id,omitempty" gorm:"<-:create;size:255;not null;index:idx_search_secondary_id,priority:1,unique;"`

	CreatedAt time.Time `json:"created_at" gorm:"<-:create;default:CURRENT_TIMESTAMP;not null;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;not null;"`
}

type MetaEntity struct {
	SecureEntity

	ShowableEntity
}

func (m *MetaEntity) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (m *MetaEntity) BeforeCreate(tx *gorm.DB) error {
	if m.SecondaryId == "" {
		m.SecondaryId = generator.RandSecondaryId()
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return nil
}
