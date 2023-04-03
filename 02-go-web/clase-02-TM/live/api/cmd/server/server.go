package main

import (
	"api/cmd/server/handlers"
	"api/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	// Env
	// ...

	// Instances
	db := []*internal.Movie{
		{ID: 1, Title: "The Shawshank Redemption", Rating: 9.2},
		{ID: 2, Title: "The Godfather", Rating: 9.2},
	}
	sv := internal.NewServiceMovieLocal(db, len(db))
	ct := handlers.NewControllerMovie(sv)


	// Server
	eng := gin.Default()
	eng.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// -> movies
	mv := eng.Group("/movies")
	{
		mv.POST("", ct.CreateMovie())
	}

	// Run
	if err := eng.Run(":8080"); err != nil {
		panic(err)
	}
}