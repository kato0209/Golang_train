package main

import (
	"Golang_train/todo/app/models"
	_ "Golang_train/todo/app/models"
	_ "Golang_train/todo/config"
)

func main() {
	u, _ := models.GetUser(1)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
}
