package api

import (
	"github.com/gin-gonic/gin"
	"thousand-hands-server/app/http/controllers"
	"thousand-hands-server/models"
	"thousand-hands-server/models/user"
)

type AboutController struct {
	controllers.BaseController
}



func (ac *AboutController) Message(c *gin.Context) {
	var users []user.User
	models.GetDB().Find(&users)
	c.JSON(200, gin.H{
		"message": "i am a message , from gin",
		"users":   users,
	})
}

