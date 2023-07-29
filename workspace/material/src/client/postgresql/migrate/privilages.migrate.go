package migrate

import (
	"gorm.io/gorm"
	"pensiel.com/domain/src/data/privilages"
	"pensiel.com/material/src/client/postgresql/executor"
	"pensiel.com/material/src/client/postgresql/options"
)

type privilagesmigrator struct {
	*gorm.DB
}

func Privilage(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&privilagesmigrator{
		dbi,
	})
}

// Instance implements execution.
func (m *privilagesmigrator) Instance() *gorm.DB {
	return m.DB
}

// Table implements migrator
func (*privilagesmigrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(privilages.EntityModel{}); !ok {
		if err := m.CreateTable(privilages.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*privilagesmigrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(privilages.EntityModel{}); ok {
		if err := m.DropTable(privilages.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*privilagesmigrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*privilagesmigrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*privilagesmigrator) Index(m *executor.Migrator) error {
	if ok := m.HasIndex(privilages.EntityModel{}, "idx_search_by_user"); !ok {
		if err := m.CreateIndex(privilages.EntityModel{}, options.INDEX, "idx_search_by_user", "user_id"); err != nil {
			return err
		}
	}

	return nil
}

// Seeder implements executor.Actions.
func (*privilagesmigrator) Seeder(m *executor.Migrator) error {
	return nil
}
