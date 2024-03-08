package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RootController is a controller for the root
type RootController struct {
	Controller
}

// Methods returns the methods for the root controller
func (u RootController) Methods(router *gin.RouterGroup) (endpoints, *gin.RouterGroup, []Controller) {
	return endpoints{
		GET: {
			endpoint{
				Path: "/",
				Function: func(c *gin.Context) {
					c.HTML(http.StatusOK, "root/index", gin.H{
						"content": "This is an index page...",
					})
				},
			},
		},
	}, router, nil
}
