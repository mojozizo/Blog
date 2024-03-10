package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserFetch(c *gin.Context) {

	user_id := c.Param("id")

	// var req models.User

	fmt.Println("User iD", user_id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Inside Fetch",
	})
}
