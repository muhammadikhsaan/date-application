package executor

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
	"pensiel.com/material/src/client/postgresql/options"
)

// Gorm Entity Interface
type entity interface {
	TableName() string
}

// Actions Helper
type Actions interface {
	Instance() *gorm.DB
	Table(m *Migrator) error
	Drop(m *Migrator) error
	Column(m *Migrator) error
	Index(m *Migrator) error
	Constraint(m *Migrator) error
	Seeder(m *Migrator) error
}

type Prosecution interface {
	Executor() error
	Reset() error
	Seeder() error
}

type prosecution struct {
	action Actions
	mt     *Migrator
}

func NewProsecution(m Actions) *prosecution {
	dbi := m.Instance()

	mt := &Migrator{
		dbi:      dbi,
		Migrator: dbi.Migrator(),
	}

	return &prosecution{
		mt:     mt,
		action: m,
	}
}

func (p *prosecution) Executor() error {
	if err := p.action.Table(p.mt); err != nil {
		return err
	}

	if err := p.action.Column(p.mt); err != nil {
		return err
	}

	if err := p.action.Index(p.mt); err != nil {
		return err
	}

	if err := p.action.Constraint(p.mt); err != nil {
		return err
	}

	return nil
}

func (p *prosecution) Seeder() error {
	if err := p.action.Seeder(p.mt); err != nil {
		return err
	}

	return nil
}

func (p *prosecution) Reset() error {
	if err := p.action.Drop(p.mt); err != nil {
		return err
	}

	return nil
}

// Migrator Helper
type Migrator struct {
	dbi *gorm.DB
	gorm.Migrator
}

func (m *Migrator) CreateIndex(dst any, types options.Types, name string, column ...string) error {
	table := dst.(entity).TableName()

	tp := options.INDEXESVALUE[types]

	query := fmt.Sprintf("CREATE %s %s ON %s (%s)", tp, name, table, strings.Join(column, ","))

	return m.dbi.Exec(query).Error
}

func (m *Migrator) Insert(dst any) error {
	return m.dbi.Create(dst).Error
}
