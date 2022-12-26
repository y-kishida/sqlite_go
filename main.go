package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/models"
)

func main() {
	todo, _ := models.GetTodo(1)
	fmt.Println(todo)

	todo.DeleteTodo()
	fmt.Println(todo)
}
