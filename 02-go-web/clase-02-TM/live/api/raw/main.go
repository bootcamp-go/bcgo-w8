package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Env
	// ...

	// Instances
	// ...

	// Server
	sv := gin.Default()
	sv.GET("/ping", func(c *gin.Context) {
		// c.JSON(200, gin.H{"data": "pong"})
		c.String(200, "pong")
	})

	// -> movies
	mv := sv.Group("/movies")
	mv.POST("", CreateMovie())

	// Run
	if err := sv.Run(":8080"); err != nil {
		panic(err)
	}
}

// package handlers
var Token = "123"

func CreateMovie() gin.HandlerFunc {
	type request struct {
		Title  string  `json:"title"`
		Rating float64 `json:"rating"`
	}
	
	return func(c *gin.Context) {
		// auth
		token := c.GetHeader("token")
		if token != Token {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "data": nil})
			return
		}

		// request
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"message": "invalid request", "data": nil})
			return
		}

		// process
		mv, err := SaveMovie(req.Title, req.Rating)
		if err != nil {
			if errors.Is(err, ErrMovieInvalid) {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid movie", "data": nil})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": nil})
			return
		}


		// response
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": mv})
	}
}


// package service
type Movie struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Rating float64 `json:"rating"`
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

var movies = []*Movie{}
var lastID = 0

var (
	ErrMovieInvalid = errors.New("invalid movie")
	ErrMovieInternal = errors.New("internal error")
)

func SaveMovie(title string, rating float64) (movie *Movie, err error) {
	// instance
	movie.ID = lastID + 1
	movie.Title = title
	movie.Rating = rating

	// valid
	if err = movie.Valid(); err != nil {
		err = ErrMovieInvalid
		return
	}

	// save
	movies = append(movies, movie)
	lastID++

	return
}


// package logger
// type Logger interface{
// 	Log(message string)
// }

// type loggerLocal struct {
// }
// func (l *loggerLocal) Log(message string) {
// 	fmt.Println(message)
// }

// var LoggerLocal = &loggerLocal{}