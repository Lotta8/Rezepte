package main

import (
	"recipies/pkg/controller"
)

func main() {
	handler := controller.NewHandler()
	handler.Run()
}
