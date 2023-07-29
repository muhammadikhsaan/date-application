package postgresql

import (
	"context"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pensiel.com/material/src/pensiel"
	"pensiel.com/material/src/static"
)

var (
	MODE = static.MODE
)

type Fcx func(tx context.Context) *pensiel.Error

type Connection struct {
	Conn *gorm.DB
	context.Context
}

type Client interface {
	Cnx(ctx context.Context) context.Context
	Trx(ctx context.Context, fc Fcx) *pensiel.Error
}

type client struct {
	dbi *gorm.DB
}

var (
	host     = static.DATABASE_HOST
	port     = static.DATABASE_PORT
	user     = static.DATABASE_USER
	password = static.DATABASE_PASSWORD
	dbname   = static.DATABASE_NAME
)

func NewClient() (Client, error) {
	dbi, err := connection()

	if err != nil {
		return nil, err
	}

	return &client{
		dbi: dbi,
	}, nil
}

func connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), nil)

	if err != nil {
		return nil, err
	}

	if MODE == "DEBUG" {
		db = db.Debug()
	}

	return db, nil
}

func (p *client) connection() *pensiel.Error {
	if p.dbi != nil {
		return nil
	}

	dbi, err := connection()

	if err != nil {
		return &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to get db connection",
		}
	}

	p.dbi = dbi
	return nil
}

func (p *client) Trx(ctx context.Context, fc Fcx) *pensiel.Error {

	if err := p.connection(); err != nil {
		return err
	}

	tx := p.dbi.Begin()

	err := fc(&Connection{
		Conn:    tx,
		Context: ctx,
	})

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (p *client) Cnx(ctx context.Context) context.Context {
	if err := p.connection(); err != nil {
		panic(err)
	}

	return &Connection{
		Conn:    p.dbi,
		Context: ctx,
	}
}
