package movies

import "database/sql"

const (
	// MySQL queries
	queryGetMovie = `SELECT id, title, release_year FROM movies WHERE id = ?`
)

// constructor
func NewRepositoryMySQL(db *sql.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db}
}

// RepositoryMySQL implements the Repository interface.
type RepositoryMySQL struct {
	db *sql.DB
}

func (r *RepositoryMySQL) GetMovie(id int) (mv Movie, err error) {
	// prepare
	var stmt *sql.Stmt
	stmt, err = r.db.Prepare(queryGetMovie)
	if err != nil {
		err = ErrRepoInternal
		return
	}

	// execute
	err = stmt.QueryRow(id).Scan(&mv.ID, &mv.Title, &mv.ReleaseYear)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrRepoNotFound
		} else {
			err = ErrRepoInternal
		}
		return
	}

	return
}