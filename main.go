package main

import (
	"fmt"
	"net/http"
)

var count int

func main() {
	fmt.Println("start")
	http.HandleFunc("/countup", handler)
	http.HandleFunc("/get", countHandler)
	//サーバを立ち上げている
	// curl http://localhost:8080　でアクセスできるサーバが起動する。
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><h1>Hello</h1></html>")
}
