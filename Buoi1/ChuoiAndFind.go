package main

import (
	"fmt"
	"strings"
)

// Đây là một chương trình Golang nhận đầu vào từ người dùng một chuỗi và
// kiểm tra xem chuỗi đó bắt đầu bằng ký tự 'i', kết thúc bằng ký tự 'n' và
// chứa ký tự 'a'. Nếu thoả mãn điều kiện trên, chương trình sẽ in ra "Found!" ngược lại nếu không thoả mãn điều kiện, chương trình sẽ in ra "Not Found!"
func main() {
	var input string

	fmt.Println("Enter a string:")
	fmt.Scanln(&input)

	// Chuyển đổi chuỗi về kiểu chữ thường để xử lý
	input = strings.ToLower(input)

	// Kiểm tra điều kiện và in ra kết quả
	//HasPrefix là một phương thức trong golang được sử dụng để kiểm tra xem một chuỗi có bắt đầu bằng một chuỗi con đã cho hay không.
	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
