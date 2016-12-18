package actions

import (
	"testing"
	"os"
	"reflect"
	"time"
	models "../models"
)

var objectPropertyStateSliceForTest = []models.ObjectPropertyState {
	models.ObjectPropertyState{
        ObjectId: "1",
        ObjectType: "ObjectA",
        Timestamp: time.Unix(412351252, 0),
        Property: "property1",
        Value: "value",
    },
    models.ObjectPropertyState{
        ObjectId: "2",
        ObjectType: "ObjectA",
        Timestamp: time.Unix(451232134, 0),
        Property: "property3",
        Value: "value",
    },
    models.ObjectPropertyState{
        ObjectId: "1",
        ObjectType: "ObjectB",
        Timestamp: time.Unix(456662343, 0),
        Property: "property1",
        Value: "value",
    },
    models.ObjectPropertyState{
        ObjectId: "1",
        ObjectType: "ObjectA",
        Timestamp: time.Unix(467765765, 0),
        Property: "property2",
        Value: "value",
    },
    models.ObjectPropertyState{
        ObjectId: "2",
        ObjectType: "ObjectA",
        Timestamp: time.Unix(451232123, 0),
        Property: "property2",
        Value: "value",
    },
    models.ObjectPropertyState{
        ObjectId: "1",
        ObjectType: "ObjectA",
        Timestamp: time.Unix(427765765, 0),
        Property: "property1",
        Value: "altered2 value",
    },
}

func TestParseCsv(t *testing.T) {
	f, err := os.Open("../test_resources/test.csv")
    if err != nil {
        t.Error("Error opening the file")
    }
    objectPropertyStateSlice := parseCsv(f)
    if !reflect.DeepEqual(objectPropertyStateSlice, objectPropertyStateSliceForTest) {
    	t.Error("The output slice does not match the expected slice")
    }
}
