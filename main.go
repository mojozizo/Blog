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
	connections.SQLDbManager()
	connections.PostgreSQLDbManager()
}

func main() {
	fmt.Println("Hello from the otherside")

	r := gin.Default()

	connections.DB.AutoMigrate(&models.User{})
	// connections.DB.AutoMigrate(&models.Blog{})

	routes.AddUserRoutes(r)
	routes.AddBlogRoutes(r)

	r.Run()
}
