package main

import (
	"fmt"
	"udemy-golang/go_crud_app/app/models"
	"udemy-golang/go_crud_app/config"
)

func main() {
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(models.Db)
}
