package user

import (
	"blog/connections"
	"blog/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserFetch(c *gin.Context) {
	var req models.User
	user_id := c.Param("id")

	res := connections.DB.First(&req, "id = ?", user_id)
	if res.Error != nil {
		fmt.Println("Error in fetching a user row", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User fetched ": req,
	})
}

func UserFetchAll(c *gin.Context) {
	var req models.User

	res := connections.DB.Find(&req)
	if res.Error != nil {
		fmt.Println("Error in fetching users", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"User fetched ": req,
	})
}
