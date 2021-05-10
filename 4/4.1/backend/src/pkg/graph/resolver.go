package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"database/sql"

	"github.com/go-playground/validator"
)

// Resolver ...
type Resolver struct {
	db *sql.DB
	v  *validator.Validate
}

// NewResolver ...
func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		db: db,
		v:  validator.New(),
	}
}
