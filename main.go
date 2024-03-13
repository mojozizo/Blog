package main

import (
	"blog/connections"
	"blog/models"
	"blog/routes"
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

	routes.AddUserRoutes(r)
	routes.AddBlogRoutes(r)

	r.Run()
}
