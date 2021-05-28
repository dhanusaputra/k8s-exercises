package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/dhanusaputra/k8s-exercises/main-project/backend/pkg/graph/generated"
	"github.com/dhanusaputra/k8s-exercises/main-project/backend/pkg/graph/model"
	"github.com/mitchellh/mapstructure"
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

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string, modifications map[string]interface{}) (*model.Todo, error) {
	rows, err := r.db.Query("SELECT id, text, done FROM todo WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}
		return nil, fmt.Errorf("cannot find ID, ID: %s", id)
	}

	updatedTodo := &model.Todo{}
	if err := rows.Scan(&updatedTodo.ID, &updatedTodo.Text, &updatedTodo.Done); err != nil {
		return nil, err
	}

	if rows.Next() {
		return nil, fmt.Errorf("find multiple rows, ID: %s", id)
	}

	if err := mapstructure.Decode(modifications, updatedTodo); err != nil {
		return nil, err
	}

	v := &vUpdateTodo{
		ID:   updatedTodo.ID,
		Text: updatedTodo.Text,
	}
	if err := r.v.Struct(v); err != nil {
		return nil, err
	}

	stmt, err := r.db.Prepare("UPDATE todo SET text=$1, done=$2 WHERE id=$3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(updatedTodo.Text, updatedTodo.Done, id)
	if err != nil {
		return nil, err
	}

	rowsNum, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsNum == 0 {
		return nil, fmt.Errorf("cannot find ID, ID: %s", id)
	}

	todo := &model.Todo{
		ID:   updatedTodo.ID,
		Text: updatedTodo.Text,
		Done: updatedTodo.Done,
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
