package main

import (
	"fmt"
)

func checkUserNameExist(username string) (bool, error) {
	if username == "ivanary" {
		return true, ErrUserNameExist{UserName: username}
	}

	return false, nil
}

func main() {
	if _, err := checkUserNameExist("foo"); err == nil {
		fmt.Println("foo not exist")
	}

	if _, err := checkUserNameExist("ivanary"); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := checkUserNameExist("ivanary"); err != nil {
		if ok := IsErrUserNameExist(err); ok {
			fmt.Println("user ivanary already exist.")
		}
	}
}
