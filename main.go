package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/models"
)

func main() {
	todos, _ := models.GetTodos()
	for i := range todos {
		fmt.Println(todos[i])
	}
}
