package blog

import (
	"blog/connections"
	"blog/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlogDelete(c *gin.Context) {
	var req models.Blog
	user_id := c.Param("id")

	res := connections.DB.Where("id = ?", user_id).Delete(&req)
	if res.Error != nil {
		fmt.Println("Error in deleting a user row", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User Deleted": user_id,
	})

}
