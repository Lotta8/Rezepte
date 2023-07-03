package main

import (
	"recipies/pkg/controller"
)

func main() { //main function erstellt uns ein Objekt den Controller
	handler := controller.New() //Von cntroller pkg rufen wir func new auf. Diese erzeugt uns einen Handler.
	handler.Run()               // auf Handler sagen wir Start, ab dann werden alle Services gestartet auf einen gewissen Port, URL und Json.
}
