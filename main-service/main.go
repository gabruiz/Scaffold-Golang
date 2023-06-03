package main

import (
	"fmt"

	controllername "example.com/main-service/controller"
)

func main() {
	fmt.Println("Hello world")
	controllername.MainService()
}
