package main

import (
	"github.com/nao11aihara/product-admin-api/app/controllers"
)

func main() {
	controllers.SetRouter()
	controllers.StartServer()
}