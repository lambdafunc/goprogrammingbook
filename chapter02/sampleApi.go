package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveApi)
	http.HandleFunc("/awesome", serveApi1)
	http.ListenAndServe(":5000", nil)

}

func serveApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func serveApi1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go is awesome")
}
