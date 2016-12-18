package actions 

import(
    "fmt"
    "net/http"
    "os"
    "encoding/csv"
    db "../db"
    models "../models"
)

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

    objectPropertyStateSlice := make([]models.ObjectPropertyState, 0)
    for i := 1; i < len(records); i++ {
        objectPropertyStateSlice = append(objectPropertyStateSlice, models.CreateObjectPropertyState(records[i])...)
    }
    fmt.Println(objectPropertyStateSlice)
    db.FlushTable()
    db.InsertIntoDb(objectPropertyStateSlice)
    return records
}
