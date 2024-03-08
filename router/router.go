package router

import (
	"html/template"
	"strings"

	"github.com/HatAndBread/presentable/controllers"
	"github.com/gin-gonic/gin"
)

// Init starts the gin router
func Init() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.LoadHTMLGlob("./views/**/*")
	router.Static("/assets", "./assets")

	baseGroup := router.Group("/")

	createRoutes(baseGroup, controllers.Controllers)

	router.Run(":3000")
}

func createRoutes(router *gin.RouterGroup, conts []controllers.Controller) {
	for _, c := range conts {
		methods, group, children := c.Methods(router)
		for verb, methods := range methods {
			for _, method := range methods {
				switch verb {
				case controllers.GET:
					group.GET(method.Path, method.Function)
				case controllers.POST:
					group.POST(method.Path, method.Function)
				case controllers.PUT:
					group.PUT(method.Path, method.Function)
				case controllers.DELETE:
					group.DELETE(method.Path, method.Function)
				default:
					panic("INVALID HTTP VERB")
				}
			}
		}
		if children != nil {
			createRoutes(group, children)
		}
	}
}
