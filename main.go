package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "world pj\n")
}
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// message := "Call Package:"
	fmt.Println("Running main()")

	fs := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.Handle("/", fs)

	http.ListenAndServe(":8090", nil)

}
