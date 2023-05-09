package movies

import "github.com/stretchr/testify/mock"

// ------------------------------------------------------------
// constructor
func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

// RepositoryMock is a movie repository mock.
type RepositoryMock struct {
	GetMovieMock struct {mv Movie; err error}
	GetMovieSpy  bool
}

func (rp *RepositoryMock) GetMovie(id int) (mv Movie, err error) {
	rp.GetMovieSpy = true
	
	mv = rp.GetMovieMock.mv
	err = rp.GetMovieMock.err
	return
}


// ------------------------------------------------------------
// constructor
func NewRepositoryMockTestify() *RepositoryMockTestify {
	return &RepositoryMockTestify{}
}

// RepositoryMockTestify is a movie repository mock.
type RepositoryMockTestify struct {
	mock.Mock
}

func (rp *RepositoryMockTestify) GetMovie(id int) (mv Movie, err error) {
	// input
	// -> for expectations (to check if the input is the same as the expected)
	args := rp.Called(id)

	// output
	mv = args.Get(0).(Movie)
	err = args.Error(1)
	return
}