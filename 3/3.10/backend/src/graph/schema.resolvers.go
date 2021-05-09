package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/dhanusaputra/k8s-exercises/graph/generated"
	"github.com/dhanusaputra/k8s-exercises/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	v := &vCreateTodoRequest{
		Text: input.Text,
	}
	if err := r.v.Struct(v); err != nil {
		return nil, err
	}

	newLastID := r.lastID + 1
	todo := &model.Todo{
		Text: input.Text,
		ID:   strconv.Itoa(newLastID),
	}

	r.todos = append(r.todos, todo)
	r.lastID = newLastID

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
