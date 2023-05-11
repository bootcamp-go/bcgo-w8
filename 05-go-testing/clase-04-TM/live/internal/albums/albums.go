package albums

import "errors"

// Models
// Album is an struct that represents an album
type Album struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year int64 `json:"year"`
}

// ------------------------------------------------------------------
// Interfaces
// Storage is an interface that represents an album storage
type Storage interface {
	// GetAlbums returns all albums
	GetAlbums() (al []*Album, err error)

	// Create creates an album
	Create(a *Album) (err error)
}
var (
	ErrStorageInternal = errors.New("internal storage error")
	ErrStorageAttributeNotUnique = errors.New("attribute not unique")
)

// Service is an interface that represents an album service
type Service interface {
	// GetAlbums returns all albums
	GetAlbums() (al []*Album, err error)

	// Create creates an album
	Create(a *Album) (err error)
}
var (
	ErrServiceInternal = errors.New("internal service error")
	ErrServiceAttributeNotUnique = errors.New("attribute not unique")
	ErrServiceInvalidAlbum = errors.New("invalid album")
)