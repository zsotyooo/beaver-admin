package todo

import (
	TodoModel "api/app/modules/todo/models"

	"github.com/jinzhu/copier"
)

type TodoCreateRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoUpdateRequest struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func UpdateRequestToTodo(request TodoUpdateRequest) TodoModel.Todo {
	var todo TodoModel.Todo
	copier.Copy(&todo, &request)
	return todo
}
