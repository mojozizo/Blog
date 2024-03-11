package user

import (
	"blog/connections"
	"blog/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {
	var req models.User
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Handle error")
		return
	}
	user_id := c.Param("id")

	res := connections.DB.Model(&models.User{}).Where("id = ?", user_id).Updates(&req)
	if res.Error != nil {
		fmt.Println("Error in updating a user row", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User updated": user_id,
	})
}
