package models

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"log"
	"thousand-hands-server/config"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	// GORM 的 MSYQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

// DB gorm.DB 对象
var db *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	switch config.GetString("database.default") {
	case "sqlite":
		db, err = connectSqlLiteDB()
		break
	case "mysql":
		db, err = connectMysqlDB()
		break
	default:
		db, err = connectSqlLiteDB()
		break
	}

	log.Println(err)
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	return db
}

func GetDB() *gorm.DB {
	return db
}
func connectMysqlDB() (*gorm.DB, error) {
	// 初始化 MySQL 连接信息
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	var level gormlogger.LogLevel

	if config.GetBool("app.debug") {
		// 读取不到数据也会显示
		level = gormlogger.Warn
	} else {
		// 只有错误才会显示
		level = gormlogger.Error
	}

	// 准备数据库连接池
	return gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
}
func connectSqlLiteDB() (*gorm.DB, error) {
// ":memory:"
	return gorm.Open(sqlite.Open("./tmp/tmp.db"), &gorm.Config{})

}
