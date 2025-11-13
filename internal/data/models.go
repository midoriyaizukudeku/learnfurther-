package data

import (
	"database/sql"
	"errors"
)

var (
	ERRRecodrdNotFound = errors.New("RECORD NOT FOUND")
	ERREditConfilct    = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
