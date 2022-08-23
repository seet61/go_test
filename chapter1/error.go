package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Введите хотя бы 1 float число")
		os.Exit(1)
	}

	args := os.Args
	var err error = errors.New("Ошибка")
	k := 1
	var n float64

	for err != nil {
		if k >= len(args) {
			fmt.Println("Ни один из аргументов не является числом")
			return
		}
		n, err = strconv.ParseFloat(args[k], 64)
		k++
	}

	min, max := n, n

	for i := 2; i < len(args); i++ {
		n, err = strconv.ParseFloat(args[i], 64)

		if err == nil {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}
	}

	fmt.Println("min:", min)
	fmt.Println("max:", max)
}
