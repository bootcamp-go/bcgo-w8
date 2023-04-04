package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "pong",
		})
	})

	if err = router.Run(":1897"); err != nil {
		panic(err)
	}
}
