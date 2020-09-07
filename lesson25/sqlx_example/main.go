package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int            `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

var DB *sqlx.DB

func initDB() error {
	var err error
	dsn := "root:123456@tcp(localhost:3306)/go_test"
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect mysql failed,error:", err)
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
	return nil
}

func testGet() {
	var user User
	sqlstr := "select * from user where id=?"
	err := DB.Get(&user, sqlstr, 2)
	if err != nil {
		fmt.Println("get failed,error:", err)
		return
	}
	fmt.Println(user)
}

func testSelect() {
	var users []User
	sqlstr := "select * from user where id>?"
	err := DB.Select(&users, sqlstr, 0)
	if err != nil {
		fmt.Println("select failed,error:", err)
		return
	}
	fmt.Println(users)
}

func testUpdate() {
	sqlstr := "insert into user(name,age) value(?,?)"
	res, err := DB.Exec(sqlstr, "liuwei", 2)
	if err != nil {
		fmt.Println("insert failed,error", err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println(count, "row affect")
}

func testTrunce() {
	conn, err := DB.Begin()
	if err != nil {
		if conn != nil {
			conn.Rollback()
		}
		fmt.Println(err)
		return
	}
	sqlstr := "insert into user(name,age) value(?,?)"
	_, err = conn.Exec(sqlstr, "pengdan", 18)
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return
	}
	err = conn.Commit()
	if err != nil {
		conn.Rollback()
		fmt.Println(err)
		return
	}
}

func queryDB(name string) {
	sqlstr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
	fmt.Printf("sql:%s\n", sqlstr)
	var user []User
	err := DB.Select(&user, sqlstr)
	if err != nil {
		fmt.Printf("select failed,err:%v\n", err)
		return
	}
	for _, v := range user {
		fmt.Printf("user%#v\n", v)
	}
}

func testSqlInject() {
	queryDB("abc' or 1 = 1#")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	//testGet()
	// testSelect()
	// testUpdate()
	//testTrunce()
	testSqlInject()
}
