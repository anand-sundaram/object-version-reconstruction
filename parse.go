package main 

import(
    "fmt"
    "net/http"
    "os"
    "encoding/csv"
)

func parseCsv(w http.ResponseWriter, r *http.Request, filename string) {
    fmt.Println("method:", r.Method)
    f, err := os.Open("./uploaded/" + filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println(err)
        return
    }
    for i := 0; i < len(records); i++ {
        for j := 0; j < len(records[i]); j++ {
            fmt.Print(records[i][j])
        }
        fmt.Println()
    }
}
