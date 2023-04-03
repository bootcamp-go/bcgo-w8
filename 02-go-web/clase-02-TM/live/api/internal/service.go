package internal

import "errors"

// _____________________________________________________________
// -> model
type Movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title" validate:"required"`
	Rating float64 `json:"rating" validate:"gte=0,lte=10"`
}
func (m *Movie) Valid() error {
	if m.Title == "" {
		return errors.New("title is required")
	}
	if m.Rating < 0 || m.Rating > 10 {
		return errors.New("rating must be between 0 and 10")
	}
	return nil
}

// _____________________________________________________________
// controller
func NewServiceMovieLocal(db []*Movie, lastID int) *ServiceMovie {
	return &ServiceMovie{
		db:     db,
		lastID: lastID,
	}
}

var (
	ErrMovieInvalid  = errors.New("invalid movie")
	ErrMovieInternal = errors.New("internal error")
)

type ServiceMovie struct {
	db     []*Movie
	lastID int
}
func (sv *ServiceMovie) Save(title string, rating float64) (movie *Movie, err error) {
	// instance
	movie = &Movie{
		ID:     sv.lastID + 1,
		Title:  title,
		Rating: rating,
	}

	// valid
	if err = movie.Valid(); err != nil {
		err = ErrMovieInvalid
		return
	}

	// save
	sv.db = append(sv.db, movie)
	sv.lastID++

	return
}