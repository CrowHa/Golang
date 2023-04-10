package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Nhập tên và địa chỉ từ người dùng
	var name, address string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter address: ")
	fmt.Scanln(&address)

	// Tạo map với key "name" và "address"
	data := map[string]string{
		"name":    name,
		"address": address,
	}

	// Tạo đối tượng JSON từ map
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// In nội dung đối tượng JSON
	fmt.Println(string(jsonData))
}
