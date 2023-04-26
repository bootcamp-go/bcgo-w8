package handlers

import (
	"api/internal/domain"
	"api/internal/movies"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// constructor
func NewControllerMovie(sv movies.Service) *ControllerMovie {
	return &ControllerMovie{sv: sv}
}

// controller
type ControllerMovie struct {
	sv movies.Service
}
func (ct *ControllerMovie) GetId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// process
		mv, err := ct.sv.GetId(id)
		if err != nil {
			if errors.Is(err, movies.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": mv})
	}
}

func (ct *ControllerMovie) Create() gin.HandlerFunc {
	type request struct {
		Title  string  `json:"title" binding:"required"`
		Year   int	   `json:"year" binding:"required"`
		Genre  string  `json:"genre" binding:"required"`
		Rating *float64 `json:"rating" binding:"required"`
	}

	return func(ctx *gin.Context) {
		// request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// process
		mv := &domain.Movie{
			Title:  req.Title,
			Year:   req.Year,
			Genre:  req.Genre,
			Rating: *req.Rating,
		}
		err := ct.sv.Create(mv)
		if err != nil {
			if errors.Is(err, movies.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		ctx.JSON(http.StatusCreated, gin.H{"message": "success", "data": mv})
	}
}

func (ct *ControllerMovie) Update() gin.HandlerFunc {
	type request struct {
		Title  string  `json:"title" binding:"required"`
		Year   int	   `json:"year" binding:"required"`
		Genre  string  `json:"genre" binding:"required"`
		Rating float64 `json:"rating"`
	}

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// process
		mv := &domain.Movie{
			ID:     id,
			Title:  req.Title,
			Year:   req.Year,
			Genre:  req.Genre,
			Rating: req.Rating,
		}
		if err := ct.sv.Update(id, mv); err != nil {
			if errors.Is(err, movies.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
				return
			}
			if errors.Is(err, movies.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": mv})
	}
}

func (ct *ControllerMovie) UpdateGenre() gin.HandlerFunc {
	type request struct {
		Genre string `json:"genre" binding:"required"`
	}

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// process
		mv, err := ct.sv.UpdateGenre(id, req.Genre)
		if err != nil {
			if errors.Is(err, movies.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
				return
			}
			if errors.Is(err, movies.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": mv})
	}
}

func (ct *ControllerMovie) UpdatePartial() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// -> get movie
		mv, err := ct.sv.GetId(id)
		if err != nil {
			if errors.Is(err, movies.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		if err := json.NewDecoder(ctx.Request.Body).Decode(&mv); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		mv.ID = id

		// process
		if err := ct.sv.Update(id, mv); err != nil {
			if errors.Is(err, movies.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
				return
			}
			if errors.Is(err, movies.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": mv})
	}
}

func (ct *ControllerMovie) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// process
		if err := ct.sv.Delete(id); err != nil {
			if errors.Is(err, movies.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "movie not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		ctx.Header("Location", fmt.Sprintf("/movies/%d", id))
		ctx.JSON(http.StatusNoContent, nil)
	}
}