package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"thousand-hands-server/config"
	"thousand-hands-server/models"
	"thousand-hands-server/models/user"
	"thousand-hands-server/routes"
)

type Application struct {
	engine *gin.Engine
}

var App Application

func Initialize() Application {

	// 加载配置文件
	config.Initialize()

	// 数据库链接迁移
	db := models.ConnectDB()
	migration(db)

	// 初始化应用程序
	App = Application{}
	env := config.Get("app.env")

	var engine *gin.Engine
	switch env {
	case config.EnvProd:
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		break
	case config.EnvLocal:
		engine = gin.Default()
		break
	default:
		engine = gin.Default()
		break
	}
	engine.LoadHTMLGlob("resources/views/*")
	App.engine = engine

	routes.RegisterRoutes(engine)

	return App
}

func migration(db *gorm.DB) {
	// 自动迁移
	err := db.AutoMigrate(
		&user.User{},
	)
	log.Println(err)
	var users = []user.User{{Name: "jinzhu1", Email: "xiaoming@outlook.com"}, {Name: "jinzhu2", Email: "xiaohong@outlook.com"}, {Name: "jinzhu3", Email: "xiaobai@outlook.com"}}
	if err = db.Create(&users).Error; err != nil {
		log.Println(err)
	}

}

// 运行程序
func (app *Application) Run() {
	port := config.GetString("app.port")
	err := app.engine.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
