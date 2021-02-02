package app

import (
	"api-auth-test/controller"
	"api-auth-test/middlewares"
	"github.com/gin-gonic/contrib/static"
)

func route() {
	router.Use(static.Serve("/", static.LocalFile("./web", true)))

	//router.GET("/", controller.Index)
	router.POST("/user", controller.CreateUser)
	router.POST("/todo", middlewares.TokenAuthMiddleware(), controller.CreateTodo)
	router.POST("/login", controller.Login)
	router.POST("/logout" ,middlewares.TokenAuthMiddleware(), controller.LogOut)
}
