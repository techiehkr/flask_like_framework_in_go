package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /request\n")
	io.WriteString(w, "this is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got hello\n")
	io.WriteString(w, "<h1>helo http</h1>")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Server Issue  %s\n", err)
		os.Exit(1)
	}

}
