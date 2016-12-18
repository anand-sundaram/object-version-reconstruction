package models

import(
    "fmt"
    "time"
    "strconv"
    "encoding/json"
)

type ObjectPropertyState struct {
    ObjectId string
    ObjectType string
    Timestamp time.Time
    Property string
    Value string
}

func CreateObjectPropertyState(record []string) []ObjectPropertyState {
    propertyValueMap := stringToPropertyValueMap(record[3])
    opsSlice := make([]ObjectPropertyState, 0)
    t, err := strconv.ParseInt(record[2], 10, 64)
    fmt.Println(time.Unix(t, 0))
    if err != nil {
        fmt.Println(err)
        return nil
    }
    for k, v := range propertyValueMap {
        ops := ObjectPropertyState{
            ObjectId: record[0],
            ObjectType: record[1],
            Timestamp: time.Unix(t, 0),
            Property: k,
            Value: v,
        }
        opsSlice = append(opsSlice, ops)
    }
    fmt.Println(opsSlice)
    return opsSlice
}

func stringToPropertyValueMap(str string) map[string]string {
    var dat map[string]string
    if err := json.Unmarshal([]byte(str), &dat); err != nil {
        fmt.Println(err)
    }
    return dat
}
