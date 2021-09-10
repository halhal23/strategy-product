package database

import "database/sql"

type SqlHandler struct {
	Conn *sql.DB
}