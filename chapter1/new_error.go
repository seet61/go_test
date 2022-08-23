package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("a == b. error in function")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("returnError() ended normally")
	}

	err = returnError(10, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("returnError() ended normally")
	}

	if err.Error() == "a == b. error in function" {
		fmt.Println("!!!")
	}
}
