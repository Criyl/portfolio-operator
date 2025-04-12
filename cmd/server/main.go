package main

import (
	"carroll.codes/portfolio-operator/internal/api"
	"carroll.codes/portfolio-operator/internal/controller"
)

func main() {
	go api.MainLoop()
	controller.Main()
}
