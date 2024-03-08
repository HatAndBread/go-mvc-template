package controllers

import "github.com/gin-gonic/gin"

// Controller is an interface for controllers
type Controller interface {
	Methods(*gin.RouterGroup) (endpoints, *gin.RouterGroup, []Controller)
}

// ValidVerb is a type for valid HTTP verbs
type ValidVerb string

// Valid HTTP verbs
const (
	GET    = ValidVerb("GET")
	POST   = ValidVerb("POST")
	PUT    = ValidVerb("PUT")
	DELETE = ValidVerb("DELETE")
)

type endpoints map[ValidVerb][]endpoint

type endpoint struct {
	Path     string
	Function gin.HandlerFunc
}
