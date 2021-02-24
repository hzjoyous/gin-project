package routes

import (
	"github.com/gin-gonic/gin"
	"thousand-hands-server/app/http/controllers/web"
)

func registerWebRoutes(r *gin.Engine) {
	wc := web.WelcomeController{}

	r.GET("/", wc.Show)

}

