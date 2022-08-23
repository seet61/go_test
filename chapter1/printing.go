package main

import "fmt"

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Хорошего дня \n"
	v4 := "abc"

	fmt.Println("Print")
	fmt.Print(v1, v2, v3, v4)
	fmt.Println()
	fmt.Println("Println")
	fmt.Println(v1, v2, v3, v4)
	fmt.Println("Print with spaces")
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
	fmt.Println("Printf")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)
}
