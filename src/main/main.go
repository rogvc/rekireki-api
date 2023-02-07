package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	p := os.Getenv("PORT")
	if p == "" {
		p = "4620"
	}

	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", p))
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
