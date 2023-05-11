package tasks

import "github.com/stretchr/testify/mock"

type MockedNotFakeDatabase struct {
	ResultOnGet *Task
	ErrOnGet    error
}

func (database *MockedNotFakeDatabase) Get(id string) (result *Task, err error) {
	return database.ResultOnGet, database.ErrOnGet
}

type BetterNotFakeDatabaseMock struct {
	mock.Mock
}

func (database *BetterNotFakeDatabaseMock) Get(id string) (result *Task, err error) {
	args := database.Mock.Called(id)
	return args.Get(0).(*Task), args.Error(1)
}
