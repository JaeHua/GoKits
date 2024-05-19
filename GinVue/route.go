package main

import (
	"GinVue/controller"
	"GinVue/middleware"
	"github.com/gin-gonic/gin"
)

// CollectRoute 路由配置函数
func CollectRoute(r *gin.Engine) *gin.Engine {

	//跨域
	r.Use(middleware.Cors())

	//注册
	r.POST("/api/auth/register", controller.Register)

	//登陆
	r.POST("/api/auth/login", controller.Login)

	//获取个人信息
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	//添加todo
	r.POST("/api/auth/todo", middleware.AuthMiddleware(), controller.CreateTodo)

	//查询todo
	r.GET("/api/auth/todo", middleware.AuthMiddleware(), controller.GetTodo)

	//更新todo
	r.PUT("/api/auth/todo:id", middleware.AuthMiddleware(), controller.UpdateTodo)

	//删除todo
	r.DELETE("/api/auth/todo:id", middleware.AuthMiddleware(), controller.DeleteTodo)

	//发送验证码
	r.POST("api/auth/email", controller.GetValidateCode)

	//验证验证码
	r.POST("api/auth/verify", controller.ValidateEmailCode)
	return r

}
