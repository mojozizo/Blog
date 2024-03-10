package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserFetchAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Inisde Fetch all ",
	})
}
