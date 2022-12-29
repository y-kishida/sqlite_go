package controllers

import (
	"net/http"
	"udemy-golang/go_crud_app/config"
)

func MainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
