package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/dhanusaputra/k8s-exercises/graph/model"
)

// Resolver ...
type Resolver struct {
	todos  []*model.Todo
	lastID int
}
