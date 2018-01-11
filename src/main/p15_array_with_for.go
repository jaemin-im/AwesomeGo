package main

import "fmt"

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	for i, value := range a {
		fmt.Println(i, value)
	}

	for _, value := range a {
		fmt.Println(value)
	}
}
