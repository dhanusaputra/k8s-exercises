package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/dhanusaputra/k8s-exercises/graph/model"
	"github.com/go-playground/validator"
)

// Resolver ...
type Resolver struct {
	todos  []*model.Todo
	lastID int
	v      *validator.Validate
}

// NewResolver ...
func NewResolver() *Resolver {
	return &Resolver{
		v: validator.New(),
	}
}
