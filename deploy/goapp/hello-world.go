package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received message: " + r.URL.Path[1:])
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("START THE HELLO WORLD SERVER..!!")
    log.Fatal(http.ListenAndServe(":8080", nil))
}