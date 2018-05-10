package main

import (
	"net/http"
	"io"
	)

func main() {
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":8080",nil)
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer,"Hello World China!")
}


