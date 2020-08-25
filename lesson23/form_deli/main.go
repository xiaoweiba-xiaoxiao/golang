package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Fprintf(w, "load login.html failed")
			return
		}
		t.Execute(w, nil)
	} else if method == "POST" {
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Printf("username: %v\n", username)
		fmt.Printf("password: %v\n", password)
		if username == "xiaochaoren" && password == "caixiaoxiao520" {
			fmt.Fprintf(w, "%s login sucessful", username)
		} else {
			fmt.Fprintf(w, "username or password error")
		}
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen server failed,err:%v\n", err)
		return
	}
}
