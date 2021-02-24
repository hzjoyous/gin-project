package routes

import (
	"github.com/gin-gonic/gin"
	"thousand-hands-server/app/http/controllers/api"
	"thousand-hands-server/app/middlewares"
)

func registerApiRoutes(r *gin.Engine) {
	ac := api.AboutController{}

	apiGroup := r.Group("api")
	apiGroup.Use(middlewares.Logger())

	{
		apiGroup.GET("/",ac.Message)
	}

}
