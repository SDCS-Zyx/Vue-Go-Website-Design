package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_driver = "root:123456@tcp(127.0.0.1:3306)/project?charset=utf8"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", db_driver)
	if err != nil {
		isOpen = false
		fmt.Println(err)
	} else {
		isOpen = true
	}
	checkErr(err)
	return isOpen, db
}

func QueryFromDB(db *sql.DB, s string) [][]sql.RawBytes {
	rows, err := db.Query(s)

	checkErr(err)
	defer rows.Close()

	cols, _ := rows.Columns()

	var result = make([][]sql.RawBytes, 0)

	for rows.Next() {
		var rowValue = make([]sql.RawBytes, len(cols))
		var rowAdr = make([]interface{}, len(cols))

		for i := range rowAdr {
			rowAdr[i] = &rowValue[i]
		}

		err = rows.Scan(rowAdr...)
		checkErr(err)
		result = append(result, rowValue)
	}

	return result
}

func InsertToDB(db *sql.DB, s string) int64 {
	res, err := db.Exec(s)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	// fmt.Println(id)
	return id
}

func UpdateToDB(db *sql.DB, s string) int64 {
	res, err := db.Exec(s)
	checkErr(err)

	aff, err := res.RowsAffected()
	checkErr(err)

	// fmt.Println(aff)
	return aff
}
