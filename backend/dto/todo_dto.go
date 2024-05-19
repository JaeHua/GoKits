package dto

import "GinVue/model"

type TodoDto struct {
	Title     string `json:"title"`
	Telephone string `json:"telephone"`
}

// ToTodoDto 将数据库ORM模型转换为API返回的DTO对象
func ToTodoDto(todo model.Todo) TodoDto {
	return TodoDto{
		Title:     todo.Title,
		Telephone: todo.Telephone,
	}
}
