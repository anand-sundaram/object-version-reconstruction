package models

import (
	"testing"
	"reflect"
	"time"
)

func TestStringToPropertyValueMap(t *testing.T) {
	propertyValue := `{"property1": "altered2 value", "property2": "value"}`
	propertyValueMap := stringToPropertyValueMap(propertyValue)
	m := make(map[string]string)
	m["property1"] = "altered2 value"
	m["property2"] = "value"
	if !reflect.DeepEqual(propertyValueMap, m) {
		t.Error("The property value map does match the expected map")
	}
}

func TestCreateObjectPropertyState(t *testing.T) {
	record := []string {"1", "ObjectA", "427765765", `{"property1": "altered2 value"}`}
	objectPropertyStateSlice := CreateObjectPropertyState(record)
	objectPropertyStateSliceForTest := []ObjectPropertyState {
		ObjectPropertyState {
			ObjectId: "1",
	        ObjectType: "ObjectA",
	        Timestamp: time.Unix(427765765, 0),
	        Property: "property1",
	        Value: "altered2 value",
        },
	}
	if !reflect.DeepEqual(objectPropertyStateSlice, objectPropertyStateSliceForTest) {
		t.Error("The output slice does not match the expected slice")
	}
}
