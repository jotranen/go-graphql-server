package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a keep alive endpoint
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
