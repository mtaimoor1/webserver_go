package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" || r.Method != "GET" {
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "HELLO from the handler\n")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Not able to parse the form err: %v\n", err)
		return
	}
	fmt.Fprintf(w, "POST request sucessfull\n")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name: %s\n", name)
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
