package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/controllers"
)

func main() {
	fmt.Println("start server")
	controllers.MainServer()
}
