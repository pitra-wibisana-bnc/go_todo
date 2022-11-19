package request

type CreateTodo struct {
	Title       string `example:"Title TODO"`
	Description string `example:"Desc TODO"`
}

type UpdateTodo struct {
	Title       string `example:"Title TODO"`
	Description string `example:"Desc TODO"`
}
