package routes

import (
	blog "blog/userblog"

	"github.com/gin-gonic/gin"
)

func AddBlogRoutes(r *gin.Engine) {
	a := r.Group("/blog")

	a.POST("/create", blog.BlogCreate)
	a.PUT("/update/:id", blog.BlogUpdate)
	a.DELETE("/delete/:id", blog.BlogDelete)
	a.GET("/fetch/:id", blog.BlogFetch)
	a.GET("/fetchall", blog.BlogFetchAll)
}
