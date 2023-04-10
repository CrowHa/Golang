package main

import (
	"fmt"
)

// Ép Kiểu
func main() {
	var x float64
	fmt.Println("Enter X here: ")
	fmt.Scanf("%f", &x)
	fmt.Println(x)
	var y int = int(x)
	fmt.Printf("X now :%d", y)
}
