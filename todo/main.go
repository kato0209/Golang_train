package main

import (
	"Golang_train/todo/app/controllers"
	_ "Golang_train/todo/app/models"
	_ "Golang_train/todo/config"
)

func main() {
	controllers.StartMainServer()
}
