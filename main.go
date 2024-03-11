package main

import (
	"blog/connections"
	"blog/models"
	"blog/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	connections.LoadEnv()
	connections.DbManager()
}

func main() {
	fmt.Println("Hello from the otherside")

	r := gin.Default()

	connections.DB.AutoMigrate(&models.User{})

	// a.Static("/static", "./templates")
	a := r.Group("/user")

	a.POST("/create", user.UserCreate)
	a.PUT("/update/:id", user.UserUpdate)
	a.DELETE("/delete/:id", user.UserDelete)
	a.GET("/fetch/:id", user.UserFetch)
	a.GET("/fetchall", user.UserFetchAll)

	r.Run()
}
