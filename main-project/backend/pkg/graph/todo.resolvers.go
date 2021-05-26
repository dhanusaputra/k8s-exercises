package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/dhanusaputra/k8s-exercises/pkg/graph/generated"
	"github.com/dhanusaputra/k8s-exercises/pkg/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	v := &vCreateTodo{
		Text: input.Text,
	}
	if err := r.v.Struct(v); err != nil {
		return nil, err
	}

	var id int
	if err := r.db.QueryRow("INSERT INTO todo(text, done) VALUES($1, $2) returning id;", input.Text, input.Done).Scan(&id); err != nil {
		return nil, err
	}

	todo := &model.Todo{
		ID:   strconv.Itoa(id),
		Text: input.Text,
		Done: input.Done,
	}

	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, input model.TodoInput) (*model.Todo, error) {
	v := &vUpdateTodo{
		ID:   id,
		Text: input.Text,
	}
	if err := r.v.Struct(v); err != nil {
		return nil, err
	}

	stmt, err := r.db.Prepare("UPDATE todo SET text=$1, done=$2 WHERE id=$3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(input.Text, input.Done, id)
	if err != nil {
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, fmt.Errorf("cannot find ID, ID: %s", id)
	}

	todo := &model.Todo{
		ID:   id,
		Text: input.Text,
		Done: input.Done,
	}

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	rows, err := r.db.Query("SELECT id, text, done FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*model.Todo{}
	for rows.Next() {
		t := &model.Todo{}
		if err := rows.Scan(&t.ID, &t.Text, &t.Done); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	sort.Slice(todos, func(i, j int) bool { return todos[i].ID < todos[j].ID })

	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
