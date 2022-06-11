package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {

}
func hellohandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)
	fmt.Println("Server Started at Port 3000")
	if err := http.ListenAndServe("3000", nil); err != nil {
		log.Fatal(err)
	}
}
