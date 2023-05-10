package tasks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRepositoryGetByID(t *testing.T) {
	t.Run("when the task exists (happy path)", func(t *testing.T) {
		// Arrange.
		taskIDToSearch := "task1"
		expectedTask := Task{
			ID:          "task1",
			Name:        "Comprar pan",
			Description: "Ir a la panadería y comprar pan",
			Completed:   true,
		}

		repository := DefaultRepository{
			Database: &MockedNotFakeDatabase{
				ResultOnGet: &expectedTask,
			},
		}

		// Act.
		obtainedTask, err := repository.GetByID(taskIDToSearch)

		// Assert.
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, obtainedTask)
	})

	t.Run("when the task doesn't exist", func(t *testing.T) {
		// Arrange.
		taskIDToSearch := "taskThatDoesntExist"
		expectedError := ErrTaskNotFound

		mock := new(BetterNotFakeDatabaseMock)
		repository := DefaultRepository{
			Database: mock,
		}

		mock.On("Get", taskIDToSearch).Return(
			&Task{},
			ErrTaskNotFound,
		)

		// Act.
		obtainedTask, err := repository.GetByID(taskIDToSearch)

		// Assert.
		mock.AssertExpectations(t)
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, Task{}, obtainedTask)
	})

	t.Run("when the database returns an error", func(t *testing.T) {
		// Arrange.
		taskIDToSearch := "taskThatDoesntExist"
		expectedError := errors.New("i'm an error, fix me plz")

		repository := DefaultRepository{
			Database: &MockedNotFakeDatabase{
				ResultOnGet: nil,
				ErrOnGet:    expectedError,
			},
		}

		// Act.
		obtainedTask, err := repository.GetByID(taskIDToSearch)

		// Assert.
		assert.EqualError(t, err, expectedError.Error())
		assert.Equal(t, Task{}, obtainedTask)
	})
}
