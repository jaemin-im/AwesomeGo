package main

import "fmt"

func main() {
	s := "Hello"
	i := 2

	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		if s == "Hello" {
			fmt.Println("Hello 2")
			break
		}
		fmt.Println(2)
	}
}
