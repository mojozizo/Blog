package blog

import (
	"blog/connections"
	"blog/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BlogCreate(c *gin.Context) {

	var req models.Blog
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Handle error")
		return
	}

	req.ID = uuid.NewString()

	res := connections.DB.Create(&req)
	if res.Error != nil {
		fmt.Println("Error in creating a user row", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	fmt.Println("Rows Affected: ", res.RowsAffected)

	c.JSON(http.StatusOK, gin.H{
		"User created": req.ID,
	})
}
