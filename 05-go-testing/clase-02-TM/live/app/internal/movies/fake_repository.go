package movies

// constructor
func NewRepositoryFake(db []Movie) *RepositoryFake {
	return &RepositoryFake{db: db}
}

// RepositoryFake implements the Repository interface.
type RepositoryFake struct {
	db []Movie
	GetMovieSpy bool
}

func (rp *RepositoryFake) GetMovie(id int) (mv Movie, err error) {
	rp.GetMovieSpy = true

	// find movie
	for _, m := range rp.db {
		if m.ID == id {
			mv = m
			return
		}
	}

	// not found
	err = ErrRepoNotFound
	return
}