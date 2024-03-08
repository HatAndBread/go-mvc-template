package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersPostsController is a controller for users
type UsersPostsController struct {
	Controller
}

// Methods returns the methods for the users controller
func (u UsersPostsController) Methods(router *gin.RouterGroup) (endpoints, *gin.RouterGroup, []Controller) {
	return endpoints{
		GET: {
			{
				Path: "/:id",
				Function: func(c *gin.Context) {
					id := c.Param("id")
					c.HTML(http.StatusOK, "users/posts/show", gin.H{
						"content": id,
					})
				},
			},
			{
				Path: "/",
				Function: func(c *gin.Context) {
					c.HTML(http.StatusOK, "users/posts/index", gin.H{
						"content": "This is the users posts index page...",
					})
				},
			},
		},
	}, router.Group("/:id/posts"), nil
}
