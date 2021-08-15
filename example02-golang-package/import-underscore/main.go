package main

import (
	"fmt"
	"os"
)

func main() {
	project := os.Getenv("GOLANG_PROJECT")
	fmt.Println(project)
}
