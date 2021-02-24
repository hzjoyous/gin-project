package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine){
	registerWebRoutes(r)
	registerApiRoutes(r)
}
