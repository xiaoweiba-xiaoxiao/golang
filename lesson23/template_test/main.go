package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type User struct {
	Name   string
	Sex    string
	Age    int
	Adress Adress
}

type Adress struct {
	City     string
	Province string
}

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		t, _ := template.ParseFiles("./index.html")
		t.Execute(w, nil)
	} else if method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		if password == "123456" {
			t, _ := template.ParseFiles("./welcom.html")
			t.Execute(w, username)
		}
	}
}

func welcom(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		var users []*User
		for i := 0; i < 10; i++ {
			user := User{
				Name: fmt.Sprintf("heihei%d", i),
				Sex:  "女",
				Age:  rand.Intn(29),
				Adress: Adress{
					City:     "chengdu",
					Province: "sichuang",
				},
			}
			users = append(users, &user)
		}
		t, _ := template.ParseFiles("./welcom.html")
		t.Execute(w, users)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/login", login)
	http.HandleFunc("/welcom", welcom)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Listen port faild, ERROR: ", err)
		return
	}
}
