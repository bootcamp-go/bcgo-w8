package albums

import "github.com/stretchr/testify/mock"

func NewStorageMock() *StorageMock {
	return &StorageMock{}
}

// ------------------------------------------------------------------
// Mock Storage
type StorageMock struct {
	mock.Mock
}

func (m *StorageMock) GetAlbums() (al []*Album, err error) {
	// input (expectations)
	args := m.Called()

	// output
	al = args.Get(0).([]*Album)
	err = args.Error(1)
	return 
}

func (m *StorageMock) Create(a *Album) (err error) {
	// input (expectations)
	args := m.Called(a)

	// output
	err = args.Error(0)
	return 
}