package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")

	fmt.Println("hi")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
