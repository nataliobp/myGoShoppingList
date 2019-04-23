package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/nataliobp/myGoShoppingList/controller"
	"github.com/nataliobp/myGoShoppingList/core"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetOutput(os.Stdout)

	container := core.NewContainer()
	container = container.Init()

	homeController := controller.HomeController{}
	userController := controller.UserController{Container: container}

	http.HandleFunc("/", homeController.Home)
	http.HandleFunc("/users", userController.ManageUser)
	http.HandleFunc("/users/", userController.ManageUser)
	http.HandleFunc("/users/email/", userController.GetUserByEmail)

	_ = http.ListenAndServe(":8001", new(core.LoggingMiddleware))
}
