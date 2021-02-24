package config

func databaseConfigSet(){
	Add("database", StrMap{
		"default":Env("DB_CONNECTION","sqlite"),
		"mysql": map[string]interface{}{

			// 数据库连接信息
			"host":     Env("DB_HOST", "127.0.0.1"),
			"port":     Env("DB_PORT", "3306"),
			"database": Env("DB_DATABASE", "goblog"),
			"username": Env("DB_USERNAME", ""),
			"password": Env("DB_PASSWORD", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": Env("DB_MAX_IDLE_CONNECTIONS", 20),
			"max_open_connections": Env("DB_MAX_OPEN_CONNECTIONS", 20),
			"max_life_seconds":     Env("DB_MAX_LIFE_SECONDS", 5*60),
		},
	})
}