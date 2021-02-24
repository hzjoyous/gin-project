package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thousand-hands-server/app/http/controllers"
)

type WelcomeController struct {
	controllers.BaseController
}

func (ac *WelcomeController) Show(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.tmpl", gin.H{
		"title": "Gin",
	})
}
