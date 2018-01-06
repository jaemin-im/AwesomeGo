package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if b, err := ioutil.ReadFile("hello.txt"); err == nil {
		fmt.Printf("%s", b)
	} else {
		fmt.Println(err)
	}
}
