package main

import "fmt"

func main() {
	i, j := 42, 1024

	p := &i
	fmt.Println(*p)

	*p = *p / 2
	fmt.Println(i)

	p = &j
	*p = *p / 2
	fmt.Println(j)
}
