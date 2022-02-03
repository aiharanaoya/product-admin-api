package main

import (
	"github.com/nao11aihara/product-admin-api/app/controllers"

	// init関数コールのため空importする
	_ "github.com/nao11aihara/product-admin-api/app/models"
)

func main() {
	controllers.SetRouter()
	controllers.StartServer()
}