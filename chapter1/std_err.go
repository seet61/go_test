package main

import (
	"io"
	"os"
)

func main() {
	my_string := ""
	arguments := os.Args
	if len(arguments) == 1 {
		my_string = "Введите 1 аргумент"
	} else {
		my_string = arguments[1]
	}
	io.WriteString(os.Stdout, "Это стандартный вывод\n")
	io.WriteString(os.Stderr, my_string)
	io.WriteString(os.Stderr, "\n")
}
