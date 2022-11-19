package views

import "time"

type CreateTodoPayload struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Title TODO"`
	Description string    `json:"description" example:"Desc TODO"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateTodoPayload struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Title TODO"`
	Description string    `json:"description" example:"Desc TODO"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetTodoPayload struct {
	ID          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Title TODO"`
	Description string    `json:"description" example:"Desc TODO"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTodoSuccessSwag struct {
	Status  int               `json:"status" example:"201"`
	Message string            `json:"message" example:"CREATE TODO SUCCESS"`
	Payload CreateTodoPayload `json:"payload"`
}

type CreateTodoFailureSwag struct {
	Status         int               `json:"status" example:"400"`
	Message        string            `json:"message" example:"CREATE TODO FAIL"`
	Error          string            `json:"message" example:"BAD Request"`
	AdditionalInfo map[string]string `json:"additional_info" example:"error:Title is required"`
}

type UpdateTodoSuccessSwag struct {
	Status  int               `json:"status" example:"200"`
	Message string            `json:"message" example:"UPDATE TODO BY ID SUCCESS"`
	Payload UpdateTodoPayload `json:"payload"`
}

type DeleteTodoSuccessSwag struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"DELETE TODO BY ID SUCCESS"`
	//Payload GetTodoPayload `json:"payload"`
}

type GetTodosSuccessSwag struct {
	Status  int              `json:"status" example:"200"`
	Message string           `json:"message" example:"GET ALL TODO SUCCESS"`
	Payload []GetTodoPayload `json:"payload"`
}

type GetTodoSuccessSwag struct {
	Status  int            `json:"status" example:"200"`
	Message string         `json:"message" example:"GET TODO BY ID SUCCESS"`
	Payload GetTodoPayload `json:"payload"`
}
