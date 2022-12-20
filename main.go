package main

import (
	"udemy-golang/go_crud_app/config"
)

func main() {
	_, err := config.NewDB()
	if err != nil {
		panic("init DB")
	}
}
