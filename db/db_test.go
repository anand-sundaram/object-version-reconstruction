package db

import (
	"testing"
	"time"
	config "../config"
	models "../models"
)

func TestDbInit(t *testing.T) {
	DbInit()
	rows, err := DBCon.Query(`SELECT * 
		FROM information_schema.tables
		WHERE table_schema = '` + config.DBName + `' 
	    AND table_name = 'object_property_state'`)
	if err != nil {
		t.Error("There was an error connecting to the database")
	}
	if !rows.Next() {
		t.Error("The table object_property_state does not exist in the database")
	}
}

func TestFlushTable(t *testing.T) {
	FlushTable()
	sqlStr := `SELECT *
		FROM object_property_state`
	rows, err := DBCon.Query(sqlStr)
	if err != nil {
		t.Error("There was an error connecting to the database")
	}
	if rows.Next() {
		t.Error("The table object_property_state does not become empty")
	}
}

func TestInsertIntoDb(t *testing.T) {
	objectPropertyStateSlice := make([]models.ObjectPropertyState, 0)
	ops := models.ObjectPropertyState{
        ObjectId: "idForTest",
        ObjectType: "typeForTest",
        Timestamp: time.Unix(1, 0),
        Property: "propertyForTest",
        Value: "valueForTest",
    }
    objectPropertyStateSlice = append(objectPropertyStateSlice, ops)
    InsertIntoDb(objectPropertyStateSlice)
 	sqlStr := `SELECT *
		FROM object_property_state
		WHERE object_type = ?
		AND object_id = ?`
	rows, err := DBCon.Query(sqlStr, "typeForTest", "idForTest")
	if err != nil {
		t.Error("There was an error connecting to the database")
	}
	if !rows.Next() {
		t.Error("Insertion into database is not working")
	}
}
