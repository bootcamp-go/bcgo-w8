package movies

import (
	"api/internal/domain"
)

// constructor
func NewRepositoryLocal(db []*domain.Movie, lastId int) Repository {
	return &repositoryLocal{db: db, lastId: lastId}
}

// controller
type repositoryLocal struct {
	db 	   []*domain.Movie
	lastId int
}
func (rp *repositoryLocal) GetId(id int) (mv *domain.Movie, err error) {
	for _, m := range rp.db {
		if m.ID == id {
			mv = m
			return
		}
	}

	err = ErrRepoNotFound
	return
}

func (rp *repositoryLocal) Create(mv *domain.Movie) (lastId int, err error) {
	// set id
	rp.lastId++
	mv.ID = rp.lastId

	// append to db
	rp.db = append(rp.db, mv)

	// return last id
	lastId = rp.lastId
	return
}

func (rp *repositoryLocal) Update(id int, mv *domain.Movie) (err error) {
	mv.ID = id

	// find movie
	for i, m := range rp.db {
		if m.ID == mv.ID {
			// update movie
			rp.db[i] = mv
			return
		}
	}

	err = ErrRepoNotFound
	return
}

func (rp *repositoryLocal) UpdateGenre(id int, genre string) (mv *domain.Movie, err error) {
	// find movie
	for i, m := range rp.db {
		if m.ID == id {
			// update movie
			rp.db[i].Genre = genre
			mv = rp.db[i]
			return
		}
	}

	err = ErrRepoNotFound
	return
}

func (rp *repositoryLocal) Delete(id int) (err error) {
	// find movie
	for i, m := range rp.db {
		if m.ID == id {
			// delete movie
			rp.db = append(rp.db[:i], rp.db[i+1:]...)
			return
		}
	}

	err = ErrRepoNotFound
	return
}