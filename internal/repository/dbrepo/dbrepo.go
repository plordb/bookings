package dbrepo

import (
	"database/sql"

	"github.com/plordb/bookings/internal/config"
	"github.com/plordb/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

/*
type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}
*/

// NewPostgresRepo
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
	}
}
