package main

import "fmt"

func main() {
	var a = 100
	var p = &a

	fmt.Println("a = ", a)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	*p = 2000
	fmt.Println("a (after) = ", a)
}