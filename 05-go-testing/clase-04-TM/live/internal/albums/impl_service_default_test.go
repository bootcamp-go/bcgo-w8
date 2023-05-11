package albums

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test ServiceDefault
func TestServiceDefault_GetAlbums(t *testing.T) {
	// succeed cases
	t.Run("get some albums", func(t *testing.T) {
		// arrange
		st := NewStorageMock()
		st.
			On("GetAlbums").
			Return([]*Album{
				{ID: 1, Title: "title 1", Author: "author 1", Year: 2001},
				{ID: 2, Title: "title 2", Author: "author 2", Year: 2002},
			}, nil)

		sv := NewServiceDefault(st)

		// act
		al, err := sv.GetAlbums()

		// assert
		assert.NoError(t, err)
		assert.Equal(t, []*Album{
			{ID: 1, Title: "title 1", Author: "author 1", Year: 2001},
			{ID: 2, Title: "title 2", Author: "author 2", Year: 2002},
		}, al)
		st.AssertExpectations(t)
	})
	t.Run("get no albums", func(t *testing.T) {
		// arrange
		st := NewStorageMock()
		st.
			On("GetAlbums").
			Return([]*Album{}, nil)

		sv := NewServiceDefault(st)

		// act
		al, err := sv.GetAlbums()

		// assert
		assert.NoError(t, err)
		assert.Equal(t, []*Album{}, al)
		st.AssertExpectations(t)				
	})

	// failure cases
	t.Run("storage error", func(t *testing.T) {
		// arrange
		st := NewStorageMock()
		st.
			On("GetAlbums").
			Return([]*Album{}, ErrStorageInternal)

		sv := NewServiceDefault(st)

		// act
		al, err := sv.GetAlbums()

		// assert
		assert.ErrorIs(t, err, ErrServiceInternal)
		assert.Empty(t, al)
	})
}

func TestServiceDefault_Create(t *testing.T) {
	// succeed cases
	t.Run("create an album", func(t *testing.T) {
		// arrange
		st := NewStorageMock()
		st.On("Create", &Album{Title: "title 1", Year: 2001}).Return(nil)
		// st.On("GetAlbums").Return([]*Album{}, nil)

		sv := NewServiceDefault(st)

		// act
		err := sv.Create(&Album{Title: "title 1", Year: 2001})

		// assert
		assert.NoError(t, err)
		st.AssertExpectations(t)
	})

	// failure cases
	// -> validation
	t.Run("invalid album - title", func(t *testing.T) {
		// arrange
		st := NewStorageMock()

		sv := NewServiceDefault(st)

		// act
		err := sv.Create(&Album{Title: "t", Year: 2001})

		// assert
		assert.ErrorIs(t, err, ErrServiceInvalidAlbum)
		assert.EqualError(t, err, "invalid album: title")
		st.AssertNotCalled(t, "Create")
		st.AssertExpectations(t)

	})
	t.Run("invalid album - year", func(t *testing.T) {
		// arrange
		st := NewStorageMock()

		sv := NewServiceDefault(st)

		// act
		err := sv.Create(&Album{Title: "title 1", Year: 1899})

		// assert
		assert.ErrorIs(t, err, ErrServiceInvalidAlbum)
		assert.EqualError(t, err, "invalid album: year")
		st.AssertNotCalled(t, "Create")
		st.AssertExpectations(t)
	})
	// -> storage
	t.Run("storage error - attribute not unique", func(t *testing.T) {
		// arrange
		st := NewStorageMock()
		st.On("Create", mock.Anything).Return(ErrStorageAttributeNotUnique)

		sv := NewServiceDefault(st)

		// act
		err := sv.Create(&Album{Title: "title 1", Year: 2001})

		// assert
		assert.ErrorIs(t, err, ErrServiceAttributeNotUnique)
		assert.EqualError(t, err, "attribute not unique")
		st.AssertExpectations(t)
	})
	t.Run("storage error - internal", func(t *testing.T) {
		// arrange
		st := NewStorageMock()
		st.On("Create", mock.Anything).Return(ErrStorageInternal)

		sv := NewServiceDefault(st)

		// act
		err := sv.Create(&Album{Title: "title 1", Year: 2001})

		// assert
		assert.ErrorIs(t, err, ErrServiceInternal)
		assert.EqualError(t, err, "internal service error")
		st.AssertExpectations(t)
	})
}

// func Case(a any) error {
// 	var err error
// 	if a != nil {
// 		err = a.(error)
// 	}

// 	return err
// }