package main

import (
	"blog/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello from the otherside")
	r := gin.Default()
	// r.Static("/static", "./templates")

	r.POST("/usercreate", user.UserCreate)
	r.PUT("/userupdate/{id}", user.UserUpdate)
	r.DELETE("/userdelete/{id}", user.UserDelete)
	r.GET("/userfetch//{id}", user.UserFetch)
	r.GET("/userfetchall", user.UserFetchAll)
	r.Run()
}
