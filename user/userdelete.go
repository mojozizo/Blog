package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDelete(c *gin.Context) {

	user_id := c.Param("id")

	if user_id == "123" {
		fmt.Println("Delete the id:", user_id)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Inside Delete",
	})

}
