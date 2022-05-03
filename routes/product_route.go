package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
}
