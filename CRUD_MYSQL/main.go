package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:hnhaf0604@tcp(127.0.0.1:3306)/testmysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println(selectAll(db))
}
func insert(db *sql.DB, Ten string, Luong int) error {
	query := "INSERT INTO BANG2(Ten, Luong) VALUES (?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Ten, Luong)
	fmt.Println("Done")
	if err != nil {
		return err
	}
	return nil
}

type Bang2 struct {
	Ten   string
	Luong int
}

func selectAll(db *sql.DB) ([]Bang2, error) {
	query := "SELECT * FROM Bang2"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bang2s []Bang2
	for rows.Next() {
		var bang2 Bang2
		err := rows.Scan(&bang2.Ten, &bang2.Luong)
		if err != nil {
			return nil, err
		}
		fmt.Println(bang2)
		bang2s = append(bang2s, bang2)
	}
	return bang2s, nil
}

func selectById(db *sql.DB, Luong int) (Bang2, error) {
	query := "SELECT * FROM Bang2 WHERE Luong = ?"
	row := db.QueryRow(query, Luong)

	var bang2 Bang2
	err := row.Scan(&bang2.Luong)
	if err != nil {
		return Bang2{}, err
	}
	return bang2, nil
}
func update(db *sql.DB, Ten string, Luong int) error {
	query := "UPDATE Bang2 SET Luong = ?,WHERE Ten = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Luong, Ten)
	if err != nil {
		return err
	}
	return nil
}
