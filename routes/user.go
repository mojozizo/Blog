package routes

import (
	"blog/user"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.Engine) {
	a := r.Group("/user")

	a.POST("/create", user.UserCreate)
	a.PUT("/update/:id", user.UserUpdate)
	a.DELETE("/delete/:id", user.UserDelete)
	a.GET("/fetch/:id", user.UserFetch)
	a.GET("/fetchall", user.UserFetchAll)
}
