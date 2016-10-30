package main

import (
    "log"
    "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
    http.HandleFunc("/", sayHello)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}