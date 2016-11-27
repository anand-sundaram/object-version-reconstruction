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
    objectId string
    objectType string
    timestamp time.Time
    property string
    value string
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
    for i := 0; i < len(records); i++ {
        createObjectPropertyState(records[i])
        for j := 0; j < len(records[i]); j++ {
            fmt.Print(records[i][j] + " ")
        }
        fmt.Println()
    }
    return records
}

func createObjectPropertyState(record []string) []objectPropertyState {
    propertyValueMap := stringToPropertyValueMap(record[3])
    opsArray := make([]objectPropertyState, 0)
    t, err := strconv.ParseInt(record[2], 10, 64)
    fmt.Println(time.Unix(t, 0))
    if err != nil {
        fmt.Println(err)
        return nil
    }
    for k, v := range propertyValueMap {
        ops := objectPropertyState{
            objectId: record[0],
            objectType: record[1],
            timestamp: time.Unix(t, 0),
            property: k,
            value: v,
        }
        opsArray = append(opsArray, ops)
    }
    fmt.Println(opsArray)
    return opsArray
}

func stringToPropertyValueMap(str string) map[string]string {
    var dat map[string]string
    if err := json.Unmarshal([]byte(str), &dat); err != nil {
        fmt.Println(err)
    }
    return dat
}
