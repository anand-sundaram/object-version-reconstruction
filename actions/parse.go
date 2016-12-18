package actions 

import(
    "fmt"
    "net/http"
    "os"
    "encoding/csv"
    db "../db"
    models "../models"
)

func parse(w http.ResponseWriter, r *http.Request, filename string) {
    fmt.Println("method:", r.Method)
    f, err := os.Open("." + pathSeparator + uploadFolderName + pathSeparator + filename)
    if err != nil {
        fmt.Println(err)
    }
    objectPropertyStateSlice := parseCsv(f)
    fmt.Println(objectPropertyStateSlice)
    db.FlushTable()
    db.InsertIntoDb(objectPropertyStateSlice)
}

func parseCsv(f *os.File) []models.ObjectPropertyState {
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
    return objectPropertyStateSlice
}