package movies

import "errors"

// ------------------------------------------------------------
// Models
// Movie represents a movie.
type Movie struct {
	ID          int	   `json:"id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
}

// ------------------------------------------------------------
// Repository is the interface that wraps the basic methods for a movie repository.
type Repository interface {
	// GetMovie returns a movie by its ID.
	GetMovie(id int) (mv Movie, err error)
}
var (
	ErrRepoInternal = errors.New("internal error")
	ErrRepoNotFound = errors.New("movie not found")
)

// ------------------------------------------------------------
// Service is the interface that wraps the basic methods for a movie service.
type Service interface {
	// GetMovie returns a movie by its ID.
	GetMovie(id int) (mv Movie, err error)
}
var (
	ErrServiceInternal = errors.New("internal error")
	ErrServiceNotFound = errors.New("movie not found")
)