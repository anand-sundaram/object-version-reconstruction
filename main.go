package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    db "./db"
    actions "./actions"
)

func main() {
	db.DbInit()
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", actions.Upload)
    router.HandleFunc("/display", actions.Display)
    router.HandleFunc("/all", db.All)
    router.HandleFunc("/type/{type}", db.ObjectType)
    router.HandleFunc("/type/{type}/id/{id}", db.ObjectTypeId)
    router.HandleFunc("/type/{type}/id/{id}/time/{time}", db.ObjectTypeIdTime)
    log.Fatal(http.ListenAndServe(":9090", router))
}
