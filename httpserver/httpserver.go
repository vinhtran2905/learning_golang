package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (handler HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello from go web wserver </h1>")
}

func main() {
	var h HelloHandler
	err := http.ListenAndServe("localhost:4000", h)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
