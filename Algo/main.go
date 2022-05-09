package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Keywords int    `json:"keywords"`
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s\n", name)

	w.Header().Set("Content-Type", "application/json")

	emp1 := Book{Id: "1", Name: "pramod", Keywords: 304}
	json.NewEncoder(w).Encode(emp1)

	emp2 := Book{Id: "2", Name: "prajwal", Keywords: 204}
	json.NewEncoder(w).Encode(emp2)

	emp3 := Book{Id: "3", Name: "pragathi", Keywords: 205}
	json.NewEncoder(w).Encode(emp3)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080 and type /form.html\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
