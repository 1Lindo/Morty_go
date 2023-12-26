package repo

import "database/sql"

type InterfaceDBRepo interface {
}

type DBRepo struct {
	db *sql.DB
}
