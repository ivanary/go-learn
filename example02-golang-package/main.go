package main

import (
	"fmt"

	"github.com/ivanary/go-learn/example02-golang-package/controller"
)

func main() {
	fmt.Println("leaning go")

	hi := controller.HelloWorld("ivanary")
	fmt.Println(hi)
}