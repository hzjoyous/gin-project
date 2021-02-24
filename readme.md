# 一个简易的gin项目文件

参考了 `laravel` ，但是放弃 laravel 的 facade 封装封装方式，只是进行了模块的组合。

```
├── app
│   ├── http
│   │   └── controllers
│   │       ├── api
│   │       │   └── about_controller.go
│   │       ├── base_controller.go     
│   │       └── web
│   │           └── welcome_controller.go
│   └── middlewares
│       └── simple_logger.go
├── bootstrap
│   └── bootstrap.go
├── config
│   ├── app.go
│   ├── config.go
│   └── database.go
├── go.mod
├── go.sum
├── main.go
├── models
│   ├── db.go
│   ├── model.go
│   └── user
│       └── user.go
├── readme.md
├── resources
│   └── views
│       └── welcome.tmpl
├── routes
│   ├── api.go
│   ├── route.go
│   └── web.go
├── tmp
│   ├── build-errors.log
│   ├── main.exe
│   ├── runner-build.exe
│   ├── runner-build.exe~
│   └── tmp.db
└── util
    └── types
        └── types.go
```

如果你想使用热加载进行开发。

[air](https://github.com/cosmtrek/air)
[realize](https://github.com/oxequa/realize)
[fresh](https://github.com/gravityblast/fresh)

```
air
realize start
fresh
```