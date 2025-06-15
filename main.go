package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
		slog.Info("Handled request", "method", c.Request.Method, "url", c.Request.URL)
	})

	slog.Info("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		slog.Error("Failed to start server", "error", err)
		return
	}
}
