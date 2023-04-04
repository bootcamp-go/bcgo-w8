package main

import (
	"api/cmd/api/handlers"
	"api/internal/domain"
	"api/internal/movies"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// instances
	// -> movies
	db := []*domain.Movie{}
	rp := movies.NewRepositoryLocal(db, 0)
	s  := movies.NewService(rp)
	ct := handlers.NewControllerMovie(s)
	
	// server
	sv := gin.Default()
	// -> ping
	sv.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	// -> movies
	movies := sv.Group("/movies")
	{
		movies.GET("/:id", ct.GetId())
		movies.POST("", ct.Create())
		movies.PUT("/:id", ct.Update())
		movies.PATCH("/:id/genre", ct.UpdateGenre())
		movies.PATCH("/:id", ct.UpdatePartial())
		movies.DELETE("/:id", ct.Delete())
	}

	// run
	if err := sv.Run(); err != nil {
		panic(err)
	}
}