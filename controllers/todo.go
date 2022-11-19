package controllers

import (
	"fmt"
	"latihan1/models"
	"latihan1/params/request"
	"latihan1/params/views"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

// @BasePath /api/v1
// GetTodos godoc
// @Summary Get All TODOS
// @Schemes
// @Description get all todos
// @Tags TODOS
// @Accept json
// @Produce json
// @Success 200 {object} views.GetTodosSuccessSwag
// @Router /todos [get]
func GetAll(c *gin.Context) {

	// Get Data
	dbEngine := c.MustGet("dbEngine").(*xorm.Engine)
	var rows_todos []models.Todo
	err := dbEngine.Find(&rows_todos)

	if err != nil {
		fmt.Println("Error sql statement: ", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	} else {

		resp := views.GeneralSuccessPayload{
			Status:  http.StatusOK,
			Message: "GET ALL TODO SUCCESS",
			Payload: rows_todos,
		}

		c.JSON(resp.Status, resp)
	}
}

// @BasePath /api/v1
// CreateTodo godoc
// @Summary Create TODO
// @Schemes
// @Description create todo
// @Tags TODOS
// @Accept json
// @Produce json
// @Param request body request.CreateTodo  true  "Request Body"
// @Success 200 {object} views.CreateTodoSuccessSwag
// @Failure	400	{object}	views.CreateTodoFailureSwag
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var req request.CreateTodo

	err := c.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	// Insert Data
	dbEngine := c.MustGet("dbEngine").(*xorm.Engine)

	row_todo := new(models.Todo)
	row_todo.Title = req.Title
	row_todo.Description = req.Description
	row_todo.CreatedAt = time.Now()

	affected, err := dbEngine.Insert(row_todo)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	} else {
		fmt.Println("affected:", affected)

		resp := views.GeneralSuccessPayload{
			Status:  http.StatusCreated,
			Message: "CREATE TODO SUCCESS",
			Payload: row_todo,
		}

		c.JSON(resp.Status, resp)
	}
}

// @BasePath /api/v1
// GetByID godoc
// @Summary Get TODO By ID
// @Schemes
// @Description get todos by ID
// @Tags TODOS
// @Accept json
// @Produce json
// @Param        id   path      int  true  "ToDo ID"
// @Success 200 {object} views.GetTodoSuccessSwag
// @Router /todos/{id} [get]
func GetByID(c *gin.Context) {

	var id = c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	ln_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error ID Conversion:", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	// Get By ID
	dbEngine := c.MustGet("dbEngine").(*xorm.Engine)
	var row_todo = models.Todo{ID: ln_id}
	has, err := dbEngine.Get(&row_todo)

	if err != nil {
		fmt.Println("Err : ", err.Error())
	} else {
		fmt.Println("has : ", has)

		if !has {
			c.AbortWithStatusJSON(http.StatusNotFound, nil)
			return
		} else {
			resp := views.GeneralSuccessPayload{
				Status:  http.StatusOK,
				Message: "GET TODO BY ID SUCCESS",
				Payload: row_todo,
			}

			c.JSON(resp.Status, resp)
		}

	}
}

// @BasePath /api/v1
// UpdateByID godoc
// @Summary Update TODO By ID
// @Schemes
// @Description update todos by ID
// @Tags TODOS
// @Accept json
// @Produce json
// @Param        id   path      int  true  "ToDo ID"
// @Param request body request.UpdateTodo  true  "Request Body"
// @Success 200 {object} views.UpdateTodoSuccessSwag
// @Router /todos/{id} [put]
func UpdateByID(c *gin.Context) {

	var id = c.Param("id")
	var req request.CreateTodo

	err := c.ShouldBindJSON(&req)

	if id == "" || err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	// Update Data
	dbEngine := c.MustGet("dbEngine").(*xorm.Engine)

	ln_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error ID Conversion:", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var row_existing_todo = models.Todo{ID: ln_id}
	has, _ := dbEngine.Get(&row_existing_todo)
	if !has {
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	} else {

		row_todo := new(models.Todo)
		row_todo.Title = req.Title
		row_todo.Description = req.Description
		row_todo.UpdatedAt = time.Now()

		affected, err := dbEngine.ID(id).Update(row_todo)

		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		} else {
			fmt.Println("affected:", affected)

			resp := views.GeneralSuccessPayload{
				Status:  http.StatusOK,
				Message: "UPDATE TODO SUCCESS",
				Payload: row_todo,
			}

			c.JSON(resp.Status, resp)
		}
	}
}

// @BasePath /api/v1
// DeleteByID godoc
// @Summary Delete TODO By ID
// @Schemes
// @Description delete todos by ID
// @Tags TODOS
// @Accept json
// @Produce json
// @Param        id   path      int  true  "ToDo ID"
// @Success 200 {object} views.DeleteTodoSuccessSwag
// @Router /todos/{id} [delete]
func DeleteByID(c *gin.Context) {
	var id = c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	// Delete Data
	dbEngine := c.MustGet("dbEngine").(*xorm.Engine)

	ln_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Error ID Conversion:", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var row_existing_todo = models.Todo{ID: ln_id}
	has, _ := dbEngine.Get(&row_existing_todo)
	if !has {
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	} else {
		affected, err := dbEngine.ID(id).Delete(&models.Todo{})

		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
			return
		} else {
			fmt.Println("affected:", affected)

			resp := views.GeneralSuccessPayload{
				Status:  http.StatusOK,
				Message: "DELETE TODO SUCCESS",
				Payload: nil,
			}

			c.JSON(resp.Status, resp)
		}
	}
}
