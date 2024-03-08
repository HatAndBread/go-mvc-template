package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersController is a controller for users
type UsersController struct {
	Controller
}

// Methods returns the methods for the users controller
func (u UsersController) Methods(router *gin.RouterGroup) (endpoints, *gin.RouterGroup, []Controller) {
	return endpoints{
		GET: {
			{
				Path: "/:id",
				Function: func(c *gin.Context) {
					id := c.Param("id")
					c.HTML(http.StatusOK, "users/show", gin.H{
						"content": id,
					})
				},
			},
		},
	}, router.Group("/users"), []Controller{UsersPostsController{}}
}
