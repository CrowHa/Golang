package utities

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:hnhaf0604@tcp(127.0.0.1:3306)/testmysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	getUser()

	// // Thiết lập các thông số cho kết nối cơ sở dữ liệu
	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(100)

	// // Khởi tạo đối tượng router của framework gin
	// router := gin.Default()

	// // Định nghĩa route GET /users/:id
	// router.GET("/bang2/:name", getUser)

	// // Khởi động server và lắng nghe các kết nối tới
	// if err := router.Run(":3000"); err != nil {
	// 	log.Fatal(err)
	// }
}
func getUser() {
	// // Lấy name từ đường dẫn
	// name := c.Param("name")

	// // Truy vấn dữ liệu từ cơ sở dữ liệu
	var luong int
	luong = 1
	err := db.QueryRow("SELECT * FROM bang2")
	if err != nil {
		fmt.Println("Lỗi Query")
		return
	}
	fmt.Println()

	// // Trả về kết quả
	// c.JSON(http.StatusOK, gin.H{"luong": luong})
}
