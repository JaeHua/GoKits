package controller

import (
	"GinVue/common"
	"GinVue/dto"
	"GinVue/model"
	"GinVue/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
)

// CreateTodo 存储todo
func CreateTodo(ctx *gin.Context) {
	DB := common.GetDB()
	//前端填写代办事项 点击提交后发送到此
	var todo = model.Todo{}
	var um model.User

	//从请求中拿出数据
	ctx.ShouldBind(&todo)

	//从上下文拿到user信息
	user, _ := ctx.Get("user")
	//断言
	um = user.(model.User)

	telephone := um.Telephone
	title := todo.Title
	status := todo.Status

	newTodo := model.Todo{
		Model:     gorm.Model{},
		Title:     title,
		Telephone: telephone,
		Status:    status,
	}
	log.Println(newTodo)
	//存入数据库
	DB.Create(&newTodo)
	response.Response(ctx, http.StatusOK, 200, gin.H{"todo": dto.ToTodoDto(newTodo)}, "todo存储成功")

}

// GetTodo 获得数据库中todo
func GetTodo(ctx *gin.Context) {
	var todoList []*model.Todo
	var um model.User

	DB := common.GetDB()
	//上下文获得user信息
	user, _ := ctx.Get("user")

	//断言
	um = user.(model.User)
	telephone := um.Telephone

	//根据用户telephone查询todo
	err := DB.Where("telephone = ?", telephone).Find(&todoList).Error

	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{"error": err.Error}, "查询错误")

	} else {
		response.Response(ctx, http.StatusOK, 200, gin.H{"todos": todoList}, "查询成功")

	}
}

// UpdateTodo 更新数据
func UpdateTodo(ctx *gin.Context) {
	//数据获取
	DB := common.GetDB()
	var todo = model.Todo{}
	id := strings.TrimPrefix(ctx.Param("id"), ":")

	//查询
	err := DB.Where("id = ?", id).First(&todo).Error
	if err != nil {

		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{"error": err.Error}, "更新错误")
		return
	}
	//状态更新
	todo.Status = !todo.Status

	//保存新数据
	err1 := DB.Save(&todo).Error

	if err1 != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{"error": err1.Error}, "更新保存错误")
		return
	} else {
		response.Response(ctx, http.StatusOK, 200, gin.H{"todos": todo}, "更新成功")
	}
}

// DeleteTodo 删除todo
func DeleteTodo(ctx *gin.Context) {
	//数据获取
	DB := common.GetDB()
	id := strings.TrimPrefix(ctx.Param("id"), ":")

	//删除对应todo
	err := DB.Where("id=?", id).Delete(model.Todo{}).Error
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{"error": err.Error}, "删除错误")
		return
	} else {
		response.Response(ctx, http.StatusOK, 200, nil, "删除成功")
	}
}
