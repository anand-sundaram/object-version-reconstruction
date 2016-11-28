package main 

import(
    "fmt"
    "net/http"
    "os"
    "encoding/csv"
    "time"
    "strconv"
    "encoding/json"
)

type objectPropertyState struct {
    ObjectId string
    ObjectType string
    Timestamp time.Time
    Property string
    Value string
}

func parseCsv(w http.ResponseWriter, r *http.Request, filename string) [][]string {
    fmt.Println("method:", r.Method)
    f, err := os.Open("." + pathSeparator + uploadFolderName + pathSeparator + filename)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer f.Close()
    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println(err)
        return nil
    }

    objectPropertyStateSlice := make([]objectPropertyState, 0)
    for i := 1; i < len(records); i++ {
        objectPropertyStateSlice = append(objectPropertyStateSlice, createObjectPropertyState(records[i])...)
    }
    fmt.Println(objectPropertyStateSlice)
    return records
}

func createObjectPropertyState(record []string) []objectPropertyState {
    propertyValueMap := stringToPropertyValueMap(record[3])
    opsSlice := make([]objectPropertyState, 0)
    t, err := strconv.ParseInt(record[2], 10, 64)
    fmt.Println(time.Unix(t, 0))
    if err != nil {
        fmt.Println(err)
        return nil
    }
    for k, v := range propertyValueMap {
        ops := objectPropertyState{
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
