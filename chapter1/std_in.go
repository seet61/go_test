package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file *os.File
	file = os.Stdin
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		fmt.Println(">>>", temp)
		if temp == "exit" {
			break
		}
	}
}
