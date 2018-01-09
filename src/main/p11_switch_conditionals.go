package main

import (
	"fmt"
)

func main() {
	i := 7
	switch {
	case i >= 5 && i < 10:
		fmt.Println("5 이상, 10 미만")
	case i >= 0 && i < 5:
		fmt.Println("0 이상, 5 미만")
	}
}
