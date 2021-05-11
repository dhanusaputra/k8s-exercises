package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/dhanusaputra/k8s-exercises/pkg/graph/generated"
	"github.com/dhanusaputra/k8s-exercises/pkg/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	v := &vCreateTodo{
		Text: input.Text,
	}
	if err := r.v.Struct(v); err != nil {
		return nil, err
	}

	var id int
	if err := r.db.QueryRow("INSERT INTO todo(text) VALUES($1) returning id;", input.Text).Scan(&id); err != nil {
		return nil, err
	}

	todo := &model.Todo{
		ID:   strconv.Itoa(id),
		Text: input.Text,
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	rows, err := r.db.Query("SELECT id, text FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*model.Todo{}
	for rows.Next() {
		t := &model.Todo{}
		if err := rows.Scan(&t.ID, &t.Text); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
