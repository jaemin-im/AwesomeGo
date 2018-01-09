package main

import "fmt"

func main() {
	i := 3
	switch i {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	}
}