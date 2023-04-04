package movies

import (
	"api/internal/domain"
	"errors"
)

// _________________________________________________________
// Repository: interface for the movies repository
type Repository interface {
	// GetId: get a movie by id
	GetId(id int) (mv *domain.Movie, err error)

	// Create: create a movie
	Create(mv *domain.Movie) (lastId int, err error)
	Update(id int, mv *domain.Movie) (err error)
	UpdateGenre(id int, genre string) (mv *domain.Movie, err error)
	Delete(id int) (err error)
}
// -> errors
var (
	ErrRepoInternal = errors.New("internal error")
	ErrRepoNotUnique = errors.New("movie already exists")
	ErrRepoNotFound = errors.New("movie not found")
)

// _________________________________________________________
// Service: interface for the movies service
type Service interface {
	// GetId: get a movie by id
	GetId(id int) (mv *domain.Movie, err error)

	// Create: create a movie
	Create(mv *domain.Movie) (err error)
	Update(id int, mv *domain.Movie) (err error)
	UpdateGenre(id int, genre string) (mv *domain.Movie, err error)
	Delete(id int) (err error)
}
// -> errors
var (
	ErrServiceInternal = errors.New("internal error")
	ErrServiceInvalid = errors.New("invalid movie")
	ErrServiceNotUnique = errors.New("movie already exists")
	ErrServiceNotFound = errors.New("movie not found")
)