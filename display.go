package main

import(
    "fmt"
    "html/template"
    "net/http"
    "io"
    "crypto/md5"
    "time"
    "strconv"
)

func display(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("frontend/display.html")
        t.Execute(w, token)
    } 
}
