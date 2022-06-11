package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Post Request Successful")
	name := r.FormValue("name")
	email := r.FormValue("email")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "Email=%s\n", email)
	fmt.Fprintf(w, "Address=%s\n", address)
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}
func main() {
	fileserver := http.FileServer(http.Dir("/static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)
	fmt.Println("Server Started at Port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
