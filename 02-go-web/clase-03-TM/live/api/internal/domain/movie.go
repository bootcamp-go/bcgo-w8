package domain

import "errors"

type Movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Genre  string  `json:"genre"`
	Rating float64 `json:"rating"`
}

// -> validate the struct
func (mv *Movie) Validate() (err error) {
	// required
	if mv.Title == "" || mv.Genre == "" {
		err = errors.New("title and genre are required")
		return
	}

	// year
	if mv.Year < 1888 {
		err = errors.New("year must be greater than 1888")
		return
	}

	// rating
	if mv.Rating < 0 || mv.Rating > 10 {
		err = errors.New("rating must be between 0 and 10")
	}

	return
}