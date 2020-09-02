package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int            `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int
}

var DB *sql.DB

func initDB() error {
	var err error
	dsn := "root:123456@(localhost:3306)/go_test"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open database failed,error:", err)
		return err
	}
	return nil
}

func testPrePareDate() {
	sqlstr := "select * from user where id>?"
	stmt, err := DB.Prepare(sqlstr)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("PrePare %s failed,errer：%s\n", sqlstr, err)
		return
	}
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("select failed,error:%v\n", err)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Age)
		fmt.Printf("%#v\n", user)
	}
}

func testPrePareInsert() {
	sqlstr := "update user set name=? where id=?"
	stmt, err := DB.Prepare(sqlstr)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Printf("PrePare %s failed,errer：%s\n", sqlstr, err)
		return
	}
	res, err := stmt.Exec("chao", 2)
	if err != nil {
		fmt.Println("update failed,error,", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("get affected faild,err:", err)
		return
	}
	fmt.Println("update sucessfull", n, "rows affected")
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	testPrePareInsert()
	testPrePareDate()
}
