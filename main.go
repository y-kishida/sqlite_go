package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/models"
)

func main() {
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testtest"
	u.CreateUser()

	user, _ := models.GetUser(2)
	user.CreateTodo("First Todo")
	fmt.Println(models.Db)
}
