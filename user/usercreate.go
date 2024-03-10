package user

import (
	"blog/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserCreate(c *gin.Context) {

	// if err := c.Request.ParseForm(); err != nil {
	// 	fmt.Printf("Form Error, err: %v", err)
	// 	return
	// }

	// fmt.Println("Post Request success")
	// name := c.PostForm("name")
	// email := c.PostForm("email")

	var req models.User
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Handle error")
		return
	}

	req.ID = uuid.NewString()

	fmt.Println("Name: ", req.FirstName)
	fmt.Println("Email: ", req.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Inside Create",
	})
}
