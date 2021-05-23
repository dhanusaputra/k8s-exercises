package graph

type vCreateTodo struct {
	Text string `json:"text" validate:"required,min=2,max=140"`
}

type vUpdateTodo struct {
	ID   string `json:"id" validate:"required"`
	Text string `json:"text" validate:"required,min=2,max=140"`
}
