package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bootcamp-go/bcgo-w8/05-go-testing/clase-03-TT/live/internal/tasks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTaskHandlerGetByID(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Arrange.
		taksIDToSearch := "task1"
		expectedHTTPStatusCode := http.StatusOK
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{
			"status": "success",
			"data": {
				"id": "task1",
				"name": "Comprar pan",
				"description": "Ir a la panadería y comprar pan",
				"completed": true
			}
		}`

		database := &tasks.DefaultNotFakeDatabase{
			Data: []tasks.Task{
				{
					ID:          "task1",
					Name:        "Comprar pan",
					Description: "Ir a la panadería y comprar pan",
					Completed:   true,
				},
				{
					ID:          "task2",
					Name:        "Comprar leche",
					Description: "Ir al supermercado y comprar leche",
					Completed:   false,
				},
			},
		}
		repository := &tasks.DefaultRepository{Database: database}
		service := &tasks.DefaultService{Repository: repository}
		handler := TasksHandler{Service: service}

		router := gin.New()
		router.GET("/tasks/:id", handler.GetByID)

		responseRecorder := httptest.NewRecorder()

		// Act.
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/tasks/"+taksIDToSearch,
			nil,
		))

		// Assert.
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})

	t.Run("when the task does not exist", func(t *testing.T) {
		// Arrange.
		taksIDToSearch := "taskThatDoesntExist"
		expectedHTTPStatusCode := http.StatusNotFound
		expectedHTTPHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{
			"status": "fail",
			"data": "task not found"
		}`

		database := &tasks.DefaultNotFakeDatabase{
			Data: []tasks.Task{
				{
					ID:          "task1",
					Name:        "Comprar pan",
					Description: "Ir a la panadería y comprar pan",
					Completed:   true,
				},
				{
					ID:          "task2",
					Name:        "Comprar leche",
					Description: "Ir al supermercado y comprar leche",
					Completed:   false,
				},
			},
		}
		repository := &tasks.DefaultRepository{Database: database}
		service := &tasks.DefaultService{Repository: repository}
		handler := TasksHandler{Service: service}

		router := gin.New()
		router.GET("/tasks/:id", handler.GetByID)

		responseRecorder := httptest.NewRecorder()

		// Act.
		router.ServeHTTP(responseRecorder, httptest.NewRequest(
			http.MethodGet,
			"/tasks/"+taksIDToSearch,
			nil,
		))

		// Assert.
		assert.Equal(t, expectedHTTPStatusCode, responseRecorder.Code)
		assert.Equal(t, expectedHTTPHeaders, responseRecorder.HeaderMap)
		assert.JSONEq(t, expectedResponse, responseRecorder.Body.String())
	})

	t.Run("when the task ID is invalid", func(t *testing.T) {
		// Arrange.

		// Act.

		// Assert.
	})

	t.Run("when the service layer returns an unexpected error", func(t *testing.T) {
		// Arrange.

		// Act.

		// Assert.
	})
}
