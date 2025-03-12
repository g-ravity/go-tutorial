package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Wrong Path", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not GET", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hey! Say 'Hello World'")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/submit" {
		http.Error(w, "Wrong Path", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not POST", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error in parsing form: %v", err)
		return
	}

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	mobile := r.FormValue("mobile")

	userData := map[string]interface{}{
		"name":   firstName + " " + lastName,
		"mobile": mobile,
	}

	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, "Error in JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/submit", formHandler)

	fmt.Println("Server started at PORT 8000!")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
