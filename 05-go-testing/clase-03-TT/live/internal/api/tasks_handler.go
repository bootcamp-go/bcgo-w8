package api

import (
	"net/http"

	"github.com/bootcamp-go/bcgo-w8/05-go-testing/clase-03-TT/live/internal/tasks"
	"github.com/gin-gonic/gin"
)

// TasksHandler is an HTTP handler layer for managing tasks in the system,
// as a RESTful API.
type TasksHandler struct {
	// Service is the bussiness logic layer for managing tasks.
	Service tasks.Service
}

// GetByID returns a task by its unique identifier, if it exists.
// If the task doesn't exist, it returns an error.
func (handler *TasksHandler) GetByID(ctx *gin.Context) {
	// Get the task ID from the URL.
	taskID := ctx.Param("id")
	if taskID == "" {
		ctx.JSON(http.StatusBadRequest, NewFailResponse("task identifier is required"))
		return
	}

	// Try to get the task from the service.
	taskFound, err := handler.Service.GetByID(taskID)
	if err != nil {
		switch err {
		case tasks.ErrTaskNotFound:
			ctx.JSON(http.StatusNotFound, NewFailResponse("task not found"))
		default:
			ctx.JSON(http.StatusInternalServerError, NewErrorResponse("an internal error has occurred, please try again later"))
		}
		return
	}

	// Return the task.
	// !it's recommended to use another structure for the response, instead of the entity itself!
	ctx.JSON(http.StatusOK, NewSuccessResponse(taskFound))
}
