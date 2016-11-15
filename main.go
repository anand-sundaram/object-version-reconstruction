package main

import (
    "log"
    "net/http"
    "path/filepath"
    "fmt"
	"github.com/gorilla/mux"
	"strings"
)

var uploadFolderName = "uploaded"
var pathSeparator = string(filepath.Separator)

func main() {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", upload)
    router.HandleFunc("/type/{type}", objectType)
    router.HandleFunc("/type/{type}/id/{id}", objectTypeId)
    router.HandleFunc("/type/{type}/id/{id}/time/{time}", objectTypeIdTime)
    log.Fatal(http.ListenAndServe(":9090", router))
}

func objectType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, "Object Type is ", vars["type"])

	records := parseCsv(w, r, "upload.csv")

	for i := 0; i < len(records); i++ {
		if strings.Compare(records[i][1], vars["type"]) == 0 {
			for j := 0; j < len(records[i]); j++ {
	            fmt.Fprintf(w, records[i][j] + " ")
	        }
	        fmt.Fprintln(w)
		}
	}
}

func objectTypeId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Object Type is %v and Object ID is %v", vars["type"], vars["id"])
	fmt.Fprintln(w)

	records := parseCsv(w, r, "upload.csv")

	for i := 0; i < len(records); i++ {
		if strings.Compare(records[i][1], vars["type"]) == 0 && strings.Compare(records[i][0], vars["id"]) == 0 {
			for j := 0; j < len(records[i]); j++ {
	            fmt.Fprintf(w, records[i][j] + " ")
	        }
	        fmt.Fprintln(w)
		}
	}
}
