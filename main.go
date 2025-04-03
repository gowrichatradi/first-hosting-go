package main

import (
	"fmt"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, err := w.Write([]byte("Love you 😘"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
