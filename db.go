package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	config "./config"
)

var (
    // DBCon is the connection handle
    // for the database
    DBCon *sql.DB
)

func dbInit() {
	var err error
	DBCon, err = sql.Open("mysql", config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DBName)
	if err != nil {
		panic(err)
	}

	_,err = DBCon.Exec(`CREATE DATABASE IF NOT EXISTS ` + config.DBName)
	if err != nil {
		panic(err)
	}
	

	DBCon, err = sql.Open("mysql", config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/" + config.DBName)
	if err != nil {
		panic(err)
	}
	

	_,err = DBCon.Exec(`CREATE TABLE IF NOT EXISTS object_property_state(
		object_id VARCHAR(20),
		object_type VARCHAR(20),
		timestamp DATETIME,
		property VARCHAR(20),
		value VARCHAR(20)
	);`)
	
}

func insertIntoDb(objectPropertyStates []objectPropertyState) {
	fmt.Println("entered insert")

	for _, row := range objectPropertyStates {
		fmt.Println(row)
		sqlStr := `INSERT INTO object_property_state(object_id, object_type, timestamp, property, value) VALUES (?, ?, ?, ?, ?)`
	    fmt.Println("reached prepare")
		stmt, err := DBCon.Prepare(sqlStr)
		if err != nil {
			fmt.Println("error in preparing")
			panic(err)
		}

		fmt.Println("reached execution")
		_, err = stmt.Exec(row.objectId, row.objectType, row.timestamp, row.property, row.value)
		if err != nil {
			fmt.Println("error in executing")
			panic(err)
		}
	}
	
}
