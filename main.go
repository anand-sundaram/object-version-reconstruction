package main

import (
    "log"
    "net/http"
    "path/filepath"
    "github.com/gorilla/mux"
)

var uploadFolderName = "uploaded"
var pathSeparator = string(filepath.Separator)

func main() {
	dbInit()
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", upload)
    router.HandleFunc("/type/{type}", objectType)
    router.HandleFunc("/type/{type}/id/{id}", objectTypeId)
    router.HandleFunc("/type/{type}/id/{id}/time/{time}", objectTypeIdTime)
    log.Fatal(http.ListenAndServe(":9090", router))
}
