package main

import (
	"database/sql"

	"github.com/bootcamp-go/bcgo-w8/04-go-storage/clase-01-TM/live/internal/users"
	"github.com/go-sql-driver/mysql"
)

func main() {
	databaseConfig := mysql.Config{
		User:   "root",
		Passwd: "SuperSecretPassword123",
		Addr:   "127.0.0.1:3306",
		DBName: "my_database",
	}

	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer database.Close()

	if err = database.Ping(); err != nil {
		panic(err)
	}

	repository := &users.MySQLRepository{
		Database: database,
	}

	repository.GetByID(1)
}
