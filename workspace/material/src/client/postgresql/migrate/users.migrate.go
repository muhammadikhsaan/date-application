package migrate

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"pensiel.com/domain/src/data/users"
	"pensiel.com/material/src/client/postgresql/executor"
	"pensiel.com/material/src/client/postgresql/options"
	"pensiel.com/material/src/contract"
)

type usersmigrator struct {
	*gorm.DB
}

func User(dbi *gorm.DB) executor.Prosecution {
	return executor.NewProsecution(&usersmigrator{
		dbi,
	})
}

// Instance implements execution.
func (um *usersmigrator) Instance() *gorm.DB {
	return um.DB
}

// Table implements migrator
func (*usersmigrator) Table(m *executor.Migrator) error {
	if ok := m.HasTable(users.EntityModel{}); !ok {
		if err := m.CreateTable(users.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Drop implements migrator.Execution.
func (*usersmigrator) Drop(m *executor.Migrator) error {
	if ok := m.HasTable(users.EntityModel{}); ok {
		if err := m.DropTable(users.EntityModel{}); err != nil {
			return err
		}
	}

	return nil
}

// Column implements migrator.
func (*usersmigrator) Column(m *executor.Migrator) error {
	return nil
}

// Constraint implements migrator.
func (*usersmigrator) Constraint(m *executor.Migrator) error {
	return nil
}

// Index implements migrator.
func (*usersmigrator) Index(m *executor.Migrator) error {
	if ok := m.HasIndex(users.EntityModel{}, "un_email"); !ok {

		if err := m.CreateIndex(users.EntityModel{}, options.UNIQUE, "un_email", "email"); err != nil {
			return err
		}

	}

	if ok := m.HasIndex(users.EntityModel{}, "idx_username_email_search"); !ok {

		if err := m.CreateIndex(users.EntityModel{}, options.INDEX, "idx_username_email_search", "email", "username", "deleted_at"); err != nil {
			return err
		}

	}

	return nil
}

// Seeder implements executor.Actions.
func (*usersmigrator) Seeder(m *executor.Migrator) error {
	go func() {
		m.Insert(&[]users.EntityModel{
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLH",
					},
				},
				Entity: users.Entity{
					Email:    "admin@admin.id",
					Password: "useradmin",
					Username: "useradmin",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLZ",
					},
				},
				Entity: users.Entity{
					Email:    "admin1@admin.id",
					Password: "useradmin1",
					Username: "useradmin1",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLp",
					},
				},
				Entity: users.Entity{
					Email:    "admin2@admin.id",
					Password: "useradmin2",
					Username: "useradmin2",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLM",
					},
				},
				Entity: users.Entity{
					Email:    "admin3@admin.id",
					Password: "useradmin3",
					Username: "useradmin3",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLB",
					},
				},
				Entity: users.Entity{
					Email:    "admin4@admin.id",
					Password: "useradmin4",
					Username: "useradmin4",
				},
			},
			{
				MetaEntity: contract.MetaEntity{
					ShowableEntity: contract.ShowableEntity{
						SecondaryId: "aynGK823985300PsnLS",
					},
				},
				Entity: users.Entity{
					Email:    "admin5@admin.id",
					Password: "useradmin5",
					Username: "useradmin5",
				},
			},
		})
	}()

	userss := []users.EntityModel{}
	for i := 0; i < 100; i++ {
		userss = append(userss, users.EntityModel{
			Entity: users.Entity{
				Email:    fmt.Sprintf("user-%d-%d@user.co.id", i, time.Now().Unix()),
				Password: "userpassword",
				Username: fmt.Sprintf("user-%d-%d", i, time.Now().Unix()),
			},
		})
	}

	if err := m.Insert(&userss); err != nil {
		return err
	}

	return nil
}
