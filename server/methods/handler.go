package methods

import (
	"gobin/database"
)

type Methods struct {
	db *database.Database
}

func New(db *database.Database) *Methods {
	return &Methods{db}
}
