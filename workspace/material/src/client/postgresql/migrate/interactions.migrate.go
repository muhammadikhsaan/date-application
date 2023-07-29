package migrate

import (
	"gorm.io/gorm"
	"pensiel.com/domain/src/data/interactions"
	"pensiel.com/material/src/client/postgresql/executor"
	"pensiel.com/material/src/client/postgresql/options"
)

type interactionsmigrator struct {
	*gorm.DB
}

func Interaction(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&interactionsmigrator{
		dbi,
	})
}

// Instance implements execution.
func (m *interactionsmigrator) Instance() *gorm.DB {
	return m.DB
}

// Table implements migrator
func (*interactionsmigrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(interactions.EntityModel{}); !ok {
		if err := m.CreateTable(interactions.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*interactionsmigrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(interactions.EntityModel{}); ok {
		if err := m.DropTable(interactions.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*interactionsmigrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*interactionsmigrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*interactionsmigrator) Index(m *executor.Migrator) error {
	if ok := m.HasIndex(interactions.EntityModel{}, "un_owner_target_combination"); !ok {
		if err := m.CreateIndex(interactions.EntityModel{}, options.UNIQUE, "un_owner_target_combination", "owner_id", "target_id"); err != nil {
			return err
		}
	}

	if ok := m.HasIndex(interactions.EntityModel{}, "idx_owner"); !ok {
		if err := m.CreateIndex(interactions.EntityModel{}, options.INDEX, "idx_owner", "owner_id"); err != nil {
			return err
		}
	}

	if ok := m.HasIndex(interactions.EntityModel{}, "idx_target"); !ok {
		if err := m.CreateIndex(interactions.EntityModel{}, options.INDEX, "idx_target", "target_id"); err != nil {
			return err
		}
	}

	return nil
}

// Seeder implements executor.Actions.
func (*interactionsmigrator) Seeder(m *executor.Migrator) error {
	return nil
}
