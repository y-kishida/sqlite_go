package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/models"
)

func main() {
	todo, _ := models.GetTodo(1)
	fmt.Println(todo)

	todo.Content = "update content"
	todo.UpdateTodo()
	fmt.Println("after update", todo)
}
