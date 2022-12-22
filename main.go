package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/models"
)

func main() {
	u, _ := models.GetUser(1)
	fmt.Println(u)
}
