package graph

type vCreateTodoRequest struct {
	Text string `json:"text" validate:"required,min=2,max=140"`
}
