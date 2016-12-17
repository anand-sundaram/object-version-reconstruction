package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "encoding/json"
    "time"
    "database/sql"
)

func all(w http.ResponseWriter, r *http.Request) {

	sqlStr := `SELECT *
		FROM object_property_state`
    fmt.Println("reached prepare")
	rows, err := DBCon.Query(sqlStr)
	if err != nil {
		fmt.Println("error in querying")
		panic(err)
	}

	printJsonOutput(rows, w)
}

func objectType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sqlStr := `SELECT * 
		FROM object_property_state
		WHERE object_type = ?`
    fmt.Println("reached prepare")
	rows, err := DBCon.Query(sqlStr, vars["type"])
	if err != nil {
		fmt.Println("error in querying")
		panic(err)
	}

	printJsonOutput(rows, w)
}

func objectTypeId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sqlStr := `SELECT * 
		FROM object_property_state
		WHERE object_type = ?
		AND object_id = ?`
    fmt.Println("reached prepare")
	rows, err := DBCon.Query(sqlStr, vars["type"], vars["id"])
	if err != nil {
		fmt.Println("error in querying")
		panic(err)
	}

	printJsonOutput(rows, w)
}

func objectTypeIdTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sqlStr := `SELECT ops.object_id, ops.object_type, ops.timestamp, ops.property, ops.value
		FROM (
			SELECT object_id, object_type, property, MAX(timestamp) AS timestamp
			FROM object_property_state
			WHERE object_type = ?
			AND object_id = ?
			AND UNIX_TIMESTAMP(timestamp) < ?
			GROUP BY object_id, object_type, property
			) AS ops_2,
			object_property_state as ops
		WHERE ops_2.object_id = ops.object_id
		AND ops_2.object_type = ops.object_type
		AND ops_2.property = ops.property
		AND ops_2.timestamp = ops.timestamp`
    fmt.Println("reached prepare")
	rows, err := DBCon.Query(sqlStr, vars["type"], vars["id"], vars["time"])
	if err != nil {
		fmt.Println("error in querying")
		panic(err)
	}

	printJsonOutput(rows, w)
}

func printJsonOutput(rows *sql.Rows, w http.ResponseWriter) {
	objectPropertyStateSlice := make([]objectPropertyState, 0)
	for rows.Next() {
		var objectId string
	    var objectType string
	    var timestamp time.Time
	    var property string
	    var value string

		err := rows.Scan(&objectId, &objectType, &timestamp, &property, &value)

		if err != nil {
	        fmt.Println(err)
	        return
	    }

	    ops := objectPropertyState{
            ObjectId: objectId,
            ObjectType: objectType,
            Timestamp: timestamp,
            Property: property,
            Value: value,
        }
        objectPropertyStateSlice = append(objectPropertyStateSlice, ops)
	    
	    fmt.Println(ops)
	}
	jsonOutput, err := json.MarshalIndent(objectPropertyStateSlice, "", "\t")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Fprintln(w, string(jsonOutput))
}
