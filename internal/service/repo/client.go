package repo

import "database/sql"

type Client interface {
	InsertClientQuery(db *sql.DB) error
	GetNames(DB *sql.DB) []string
}
