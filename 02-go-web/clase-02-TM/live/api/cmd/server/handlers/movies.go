package handlers

import (
	"api/internal"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var sv = internal.NewServiceMovieLocal([]*internal.Movie{}, 0)

// constructor
func NewControllerMovie(sv *internal.ServiceMovie) *ControllerMovie {
	return &ControllerMovie{sv: sv}
}

// controller
type ControllerMovie struct {
	sv *internal.ServiceMovie
}

func (ct *ControllerMovie) CreateMovie() gin.HandlerFunc {
	type request struct {
		Title  string  `json:"title"`
		Rating float64 `json:"rating"`
	}

	return func(c *gin.Context) {
		// request
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "data": nil})
			return
		}

		// process
		mv, err := ct.sv.Save(req.Title, req.Rating)
		if err != nil {
			if errors.Is(err, internal.ErrMovieInvalid) {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid movie", "data": nil})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error", "data": nil})
			return
		}


		// response
		c.JSON(http.StatusCreated, gin.H{"message": "success", "data": mv})
	}
}