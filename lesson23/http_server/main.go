package main

import (
	"fmt"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	for key, val := range r.Form {
		fmt.Println("key:", key)
		fmt.Println("value:", val)
	}
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("listen port faild,err:", err)
		return
	}
}
