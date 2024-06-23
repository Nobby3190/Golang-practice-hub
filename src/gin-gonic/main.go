package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	server.GET("/add", func(ctx *gin.Context) {
		aStr := ctx.Query("a")
		bStr := ctx.Query("b")

		a, errA := strconv.Atoi(aStr)
		b, errB := strconv.Atoi(bStr)

		if errA != nil || errB != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid parameters",
			})
			return
		}

		sum := a + b
		ctx.JSON(http.StatusOK, gin.H{
			"Sum": sum,
		})
	})

	server.Run()
}
