package movies

import "api/internal/domain"

// constructor
func NewService(rp Repository) Service {
	return &service{rp: rp}
}

// controller
type service struct {
	rp Repository
}
func (s *service) GetId(id int) (mv *domain.Movie, err error) {
	// get movie
	mv, err = s.rp.GetId(id)
	if err != nil {
		return
	}
	if mv == nil {
		err = ErrServiceNotFound
		return
	}
	return
}

func (s *service) Create(mv *domain.Movie) (err error) {
	// validate movie
	err = mv.Validate()
	if err != nil {
		err = ErrServiceInvalid
		return
	}

	// create movie
	var lastId int
	lastId, err = s.rp.Create(mv)
	if err != nil {
		err = ErrServiceInternal
		return
	}

	mv.ID = lastId
	return
}

func (s *service) Update(id int, mv *domain.Movie) (err error) {
	// validate movie
	err = mv.Validate()
	if err != nil {
		err = ErrServiceInvalid
		return
	}

	// update movie
	err = s.rp.Update(mv.ID, mv)
	if err != nil {
		if err == ErrRepoNotFound {
			err = ErrServiceNotFound
			return
		}
		err = ErrServiceInternal
	}
	
	return
}

func (s *service) UpdateGenre(id int, genre string) (mv *domain.Movie, err error) {
	// validation genre
	if genre == "" {
		err = ErrServiceInvalid
		return
	}

	// update movie
	mv, err = s.rp.UpdateGenre(id, genre)
	if err != nil {
		if err == ErrRepoNotFound {
			err = ErrServiceNotFound
			return
		}

		err = ErrServiceInternal
	}
	return
}

func (s *service) Delete(id int) (err error) {
	// delete movie
	err = s.rp.Delete(id)
	if err != nil {
		if err == ErrRepoNotFound {
			err = ErrServiceNotFound
			return
		}

		err = ErrServiceInternal
	}
	return
}