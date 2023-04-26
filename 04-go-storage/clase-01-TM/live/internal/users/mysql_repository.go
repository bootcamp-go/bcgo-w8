package users

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	Database *sql.DB
}

func (repository *MySQLRepository) GetByID(id int) (user *User, err error) {
	// Generate the query.
	query := `
		SELECT
			id, username, email
		FROM users
		WHERE id = ?
	`

	// Execute the query.
	row := repository.Database.QueryRow(query, id)

	// Map the result.
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return
	}

	// Return the result.
	return
}

func (repository *MySQLRepository) Create(user *User) (err error) {
	// Generate the statement.
	statement, err := repository.Database.Prepare(`
		INSERT INTO users (username, email)
		VALUES (?, ?)
	`)
	if err != nil {
		return
	}
	defer statement.Close()

	// Execute the statement.
	result, err := statement.Exec(
		user.Username,
		user.Email,
	)
	if err != nil {
		// Cast to MySQL error.
		mysqlError, ok := err.(*mysql.MySQLError)
		if !ok {
			return
		}

		// Check the error code.
		switch mysqlError.Number {
		case 1062:
			err = ErrAlreadyExists
		case 1586:
			err = ErrAlreadyExists
			// TODO: Handle more errors.
		}

		return
	}

	// Map the result.
	lastID, err := result.LastInsertId()
	if err != nil {
		return
	}

	// Return the result.
	user.ID = int(lastID)
	return
}

func (repository *MySQLRepository) Update(user *User) (err error) {
	// Generate the statement.
	statement, err := repository.Database.Prepare(`
		UPDATE users
		SET username = ?, email = ?
		WHERE id = ?
	`)
	if err != nil {
		return
	}

	// Execute the statement.
	_, err = statement.Exec(
		// To update...
		user.Username,
		user.Email,

		// Where...
		user.ID,
	)
	if err != nil {
		// TODO: Cast to MySQL error.
		return
	}

	// Return the result.
	return
}

func (repository *MySQLRepository) Delete(id int) (err error) {
	// Generate the statement.
	statement, err := repository.Database.Prepare(`
		DELETE FROM users
		WHERE id = ?
	`)
	if err != nil {
		return
	}

	// Execute the statement and check if the user was deleted.
	_, err = statement.Exec(id)
	if err != nil {
		// TODO: Cast to MySQL error.
		return
	}

	return
}
