package database

import (
	"database/sql"

	"gopkg.in/reform.v1"

	_ "github.com/lib/pq"

	"github.com/datal-hub/candles/models"
	"github.com/datal-hub/candles/pkg/database/pgsql"
	"github.com/datal-hub/candles/pkg/settings"
)

// DB define interface for work with database
type DB interface {
	IsEmpty() bool
	Clear()
	Init(force bool) error
	Close() error
	SqlDB() *sql.DB

	GetOrCreateTickers(labels []string) ([]models.Ticker, error)
	Save(model reform.Record) error
}

// NewDB create new database connection accordance with mode
func NewDB() (DB, error) {
	db, err := sql.Open("postgres", settings.DB.Url)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &pgsql.PgSQL{DB: db}, nil
}
