package main

import (
	"fmt"
	"golang/controller"
	"golang/service"
)

func main() {
	fmt.Println("welcome to golang ")
	fmt.Println("Lets learn go lang ")
	service.Myservice()
	controller.ServiceController()
}
