package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Khởi tạo slice với kích thước ban đầu là 3
	nums := make([]int, 3)

	for {
		var input string
		fmt.Println("Enter an integer (X to exit): ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or X to exit.")
			continue
		}

		// Thêm số vào slice
		nums = append(nums, num)

		// Sắp xếp lại slice theo thứ tự tăng dần
		sort.Ints(nums)

		fmt.Println(nums)
	}
}
