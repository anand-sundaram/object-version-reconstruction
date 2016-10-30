package main

import (
    "log"
    "net/http"
    "path/filepath"
)

var uploadFolderName = "uploaded"
var pathSeparator = string(filepath.Separator)

func main() {
    http.HandleFunc("/", upload)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}