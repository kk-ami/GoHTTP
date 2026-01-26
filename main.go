package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"io"
)
const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr) )
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Printf("$s: got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello there!\n")
}

/*
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
} */

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	//mux.HandleFunc("/headers", headers)

	fmt.Println("Server starting on http://localhost:8090...")

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.server{
		Addr: ":3333",
		Handler: mux,
		BaseContext: func(l net.Listner) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	/*
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		panic(err)
	} */
}