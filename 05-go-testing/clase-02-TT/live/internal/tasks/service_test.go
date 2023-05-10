package tasks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit tests for the DefaultService.GetByID function.
func TestDefaultServiceGetByID(t *testing.T) {
	t.Run("when the task exists (happy path)", func(t *testing.T) {
		// Arrange.
		taskIDToSearch := "task1"
		expectedTask := Task{
			ID:          "task1",
			Name:        "Comprar pan",
			Description: "Ir a la panadería y comprar pan",
			Completed:   true,
		}

		service := DefaultService{
			Repository: &MockedRepository{
				ResultOnGet: expectedTask,
			},
		}

		// Act.
		obtainedTask, err := service.GetByID(taskIDToSearch)

		// Assert.
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, obtainedTask)
	})

	t.Run("when the task does not exist", func(t *testing.T) {
		// Arrange.
		taskIDToSearch := "taskThatDoesntExist"
		expectedError := ErrTaskNotFound

		service := DefaultService{
			Repository: &MockedRepository{
				ResultOnGet: Task{},
				ErrOnGet:    ErrTaskNotFound,
			},
		}

		// Act.
		obtainedTask, err := service.GetByID(taskIDToSearch)

		// Assert.
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, Task{}, obtainedTask)
	})
}

// Integration tests for the DefaultService.GetByID function,
// using the DefaultRepository layer:
func TestDefaultServiceDefaultRepositoryGetByID(t *testing.T) {
	t.Run("when the task exists (happy path)", func(t *testing.T) {
		// Arrange
		taskIDToSearch := "task1"
		expectedTask := Task{
			ID:          "task1",
			Name:        "Comprar pan",
			Description: "Ir a la panadería y comprar pan",
			Completed:   true,
		}

		service := DefaultService{
			Repository: &DefaultRepository{
				Database: &MockedNotFakeDatabase{
					ResultOnGet: &expectedTask,
				},
			},
		}

		// Act
		obtainedTask, err := service.GetByID(taskIDToSearch)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, obtainedTask)
	})

	t.Run("when the task does not exist", func(t *testing.T) {
		// Arrange
		taskIDToSearch := "taskThatDoesntExist"
		expectedError := ErrTaskNotFound

		service := DefaultService{
			Repository: &DefaultRepository{
				Database: &MockedNotFakeDatabase{
					ResultOnGet: nil,
					ErrOnGet:    ErrTaskNotFound,
				},
			},
		}

		// Act
		obtainedTask, err := service.GetByID(taskIDToSearch)

		// Assert
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, Task{}, obtainedTask)
	})

	t.Run("when the database returns an error", func(t *testing.T) {
		// Arrange
		taskIDToSearch := "task1"
		expectedError := errors.New("database connection refused")

		service := DefaultService{
			Repository: &DefaultRepository{
				Database: &MockedNotFakeDatabase{
					ResultOnGet: nil,
					ErrOnGet:    expectedError,
				},
			},
		}

		// Act
		obtainedTask, err := service.GetByID(taskIDToSearch)

		// Assert
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, Task{}, obtainedTask)
	})
}
