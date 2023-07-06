package repo

import "database/sql"

type Goods interface {
	Insert(db *sql.DB) error
	GetId(DB *sql.DB) string
	GetName(DB *sql.DB)
	GetAmount(DB *sql.DB) (amountCheck string)
}
