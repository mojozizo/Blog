package main

import (
	"blog/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello from the otherside")
	r := gin.Default()
	r.GET("/usercreate", user.UserCreate)
	r.Run()
}
