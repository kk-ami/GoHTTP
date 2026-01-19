package main

import (
	"fmt"
	"net/http"
	"io"
)

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello there!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/headers", headers)

	fmt.Println("Server starting on http://localhost:8090...")

	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		panic(err)
	}
}