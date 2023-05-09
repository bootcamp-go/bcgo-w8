package movies

import (
	"app/internal/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------
// stub tests
func TestService_GetMovie_StubFunc(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryStubFunc()
		rp.GetMovieStub = func(id int) (mv Movie, err error) {
			mv = Movie{
				ID:          1,
				Title:       "The Matrix",
				ReleaseYear: 1999,
			}
			return
		}

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, mv)
	})

	t.Run("fail not found", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryStubFunc()
		rp.GetMovieStub = func(id int) (mv Movie, err error) {
			err = ErrRepoNotFound
			return
		}

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.Empty(t, mv)
		assert.ErrorIs(t, err, ErrServiceNotFound)
	})
}

func TestService_GetMovie_StubReturn(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryStubReturn()
		rp.GetMovieStub.mv = Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}
		rp.GetMovieStub.err = nil

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, mv)
	})

	t.Run("fail not found", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryStubReturn()
		rp.GetMovieStub.mv = Movie{}
		rp.GetMovieStub.err = ErrRepoNotFound

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.Empty(t, mv)
		assert.ErrorIs(t, err, ErrServiceNotFound)
	})
}

func TestService_GetMovie_StubArgs(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryStubArgs()
		rp.GetMovieArgs = NewArgs(Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, nil)

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, mv)
	})

	t.Run("fail not found", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryStubArgs()
		rp.GetMovieArgs = NewArgs(Movie{}, ErrRepoNotFound)

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.Empty(t, mv)
		assert.ErrorIs(t, err, ErrServiceNotFound)
	})
}


// ------------------------------------------------------------
// mock tests
func TestService_GetMovie_MockArgs(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryMock()
		rp.GetMovieMock.mv = Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}
		rp.GetMovieMock.err = nil

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, mv)
		// -> expected calls
		assert.True(t, rp.GetMovieSpy)
	})

	t.Run("fail not found", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryMock()
		rp.GetMovieMock.mv = Movie{}
		rp.GetMovieMock.err = ErrRepoNotFound
		
		sv := NewServiceDefault(rp, lg)
		
		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.Empty(t, mv)
		assert.ErrorIs(t, err, ErrServiceNotFound)
		// -> expected calls
		assert.True(t, rp.GetMovieSpy)
	})
}

func TestService_GetMovie_Testify(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryMockTestify()
		rp.
			On("GetMovie", 1).
			Return(Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, nil)
		

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, mv)
		// -> expected calls
		rp.AssertExpectations(t)
	})

	t.Run("fail not found", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		rp := NewRepositoryMockTestify()
		rp.
			On("GetMovie", 1).
			Return(Movie{}, ErrRepoNotFound)

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.Empty(t, mv)
		assert.ErrorIs(t, err, ErrServiceNotFound)
		// -> expected calls
		rp.AssertExpectations(t)
	})
}


// ------------------------------------------------------------
// fake tests
func TestService_GetMovie_Fake(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		db := []Movie{
			{ID: 1, Title: "The Matrix", ReleaseYear: 1999},
		}
		rp := NewRepositoryFake(db)

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, Movie{ID: 1, Title: "The Matrix", ReleaseYear: 1999}, mv)
		assert.True(t, rp.GetMovieSpy)
	})

	t.Run("fail not found", func(t *testing.T) {
		// arrange
		lg := logger.NewLoggerDummy()
		db := []Movie{}
		rp := NewRepositoryFake(db)

		sv := NewServiceDefault(rp, lg)

		// act
		mv, err := sv.GetMovie(1)

		// assert
		assert.Empty(t, mv)
		assert.ErrorIs(t, err, ErrServiceNotFound)
		assert.True(t, rp.GetMovieSpy)
	})
}