package main

import (
    "log"
    "net/http"
    "path/filepath"
    "fmt"
    "github.com/gorilla/mux"
    "strings"
    "strconv"
    "encoding/json"
    "sort"
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

func objectTypeIdTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Object Type is %v and Object ID is %v and the Time is %v", vars["type"], vars["id"], vars["time"])
	fmt.Fprintln(w)

	records := parseCsv(w, r, "upload.csv")

	time, err := strconv.Atoi(vars["time"])
	if err != nil {
        fmt.Println(err)
    }

	properties := ""
	var mapTimeToChanges map[int](map[string]string)

	mapTimeToChanges = make(map[int](map[string]string))

	for i := 0; i < len(records); i++ {
		recordTimestamp, err := strconv.Atoi(records[i][2])
		if err != nil {
	        fmt.Println(err)
	    }
		if strings.Compare(records[i][1], vars["type"]) == 0 && strings.Compare(records[i][0], vars["id"]) == 0 && time >= recordTimestamp {
			properties = properties + records[i][3] + ", "
			for j := 0; j < len(records[i]); j++ {
	            fmt.Fprintf(w, records[i][j] + " ")
	        }
	        fmt.Fprintln(w)
			var dat map[string]string
            if err := json.Unmarshal([]byte(records[i][3]), &dat); err != nil {
		        fmt.Println(err)
		    }
		    mapTimeToChanges[recordTimestamp] = dat
		    fmt.Println(dat)
		    for key, value := range dat {
			    fmt.Printf("key[%s] value[%s]\n", key, value)
			}
	        fmt.Fprintln(w)
		}
	}

	var keys []int
    for k := range mapTimeToChanges {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    var objectStateMap map[string]string
    objectStateMap = make(map[string]string)

    for i := 0; i < len(keys); i++ {
		var m map[string]string
		m = make(map[string]string)
		m = mapTimeToChanges[keys[i]]
		for key, value := range m {
		    fmt.Printf("key[%s] value[%s]\n", key, value)
		    objectStateMap[key] = value
		}
    }
    fmt.Fprintln(w, objectStateMap)
}