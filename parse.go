package main 

import(
    "fmt"
    "net/http"
    "os"
    "encoding/csv"
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
    for i := 0; i < len(records); i++ {
        for j := 0; j < len(records[i]); j++ {
            fmt.Print(records[i][j] + " ")
        }
        fmt.Println()
    }
    return records
}
