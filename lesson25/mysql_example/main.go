package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB() error {
	var err error
	dsn := "root:123456@tcp(192.168.234.200:3306)/go_test"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100) //设置最大连接
	DB.SetMaxIdleConns(16)  //设置最大空闲连接
	return nil
}

type Person struct {
	Id   int            `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

func testQuery() {
	sqlstr := "select * from user where id>?"
	rows, err := DB.Query(sqlstr, 0)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Println("query faild, error", err)
		return
	}
	var persons []*Person
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			continue
		}
		persons = append(persons, &person)
	}
	for _, val := range persons {
		fmt.Println(*val)
	}
}

func testQueryRow() {
	sqlstr := "select * from user where id=?"
	result := DB.QueryRow(sqlstr, 2)

	var person Person
	err := result.Scan(&person.Id, &person.Name, &person.Age)
	if err != nil {
		fmt.Printf("scan error：%v\n", err)
		return
	}
	fmt.Println(person)
	// fmt.Println(columns)
}

func testInsert() {
	sqlstr := "insert into user(name,age) values(?,?),(?,?)"
	result, err := DB.Exec(sqlstr, "xiaoxiao", 18, "xiaochaoren", 19)
	if err != nil {
		fmt.Println("insert faild,err", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get insert id faild,err:", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get affected faild,err:", err)
		return
	}
	fmt.Println(id, "insert sucessful", n, "rows affectd")
}

func testUpdate() {
	sqlstr := "update user set name=? where id=?"
	result, err := DB.Exec(sqlstr, "wangchao", 2)
	if err != nil {
		fmt.Println("update faild,error", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get affected faild,err:", err)
		return
	}
	fmt.Println("update sucessfull", n, "rows affected")
}

func testDelete() {
	sqlstr := "delete from user where id=?"
	result, err := DB.Exec(sqlstr, 4)
	if err != nil {
		fmt.Println("delete faild,err", err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get affected faild,err:", err)
		return
	}
	fmt.Println("update sucessfull", n, "rows affected")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// testUpdate()
	//testQuery()
	// testInsert()
	testDelete()
	testQuery()
}
