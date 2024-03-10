package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Inside Update",
	})
}
