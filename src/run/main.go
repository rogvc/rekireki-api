package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if p := os.Getenv("PORT"); p == "" {
		p = "4620"
	}

	r := setupRouter()
	r.Run("localhost:4620")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	return r
}
