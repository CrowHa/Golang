package main

import (
	"fmt"
)

func main() {
	fmt.Println("he")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("Loi o day ne")

}