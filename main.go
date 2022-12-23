package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/models"
)

func main() {
	user, _ := models.GetUser(0)

	todos, _ := user.GetTodosByUser()
	for i := range todos {
		fmt.Println(todos[i])
	}
}
