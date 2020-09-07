package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var count int64

func initDB() error {
	var err error
	dsn := "root:123456@tcp(localhost:3306)/go_test"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open database failed,error", err)
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	return nil
}

func testCommit() {
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Println("trance failed,error", err)
		return
	}

	sqlstr := "insert into user(name,age) values(?,?)"
	res, err := conn.Exec(sqlstr, "laosiji", 18)
	if err != nil {
		conn.Rollback()
		fmt.Println("insert failed,error", err)
		return
	}
	n, _ := res.RowsAffected()
	count += n
	res, err = conn.Exec(sqlstr, "chenshao", 18)
	if err != nil {
		conn.Rollback()
		fmt.Println("insert failed,error", err)
		return
	}
	n, _ = res.RowsAffected()
	count += n
	err = conn.Commit()
	if err != nil {
		conn.Rollback()
		fmt.Println("commit failed,error", err)
		return
	}
	fmt.Println(count, "rows affect")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	testCommit()
}
