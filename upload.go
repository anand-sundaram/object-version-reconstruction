package main

import(
    "fmt"
    "html/template"
    "net/http"
    "io"
    "os"
    "crypto/md5"
    "time"
    "strconv"
)

func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("frontend/upload.html")
        t.Execute(w, token)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        
        os.Mkdir("." + pathSeparator + uploadFolderName, 0777)
        f, err := os.OpenFile("." + pathSeparator + uploadFolderName + pathSeparator + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
        parseCsv(w, r, handler.Filename)
        http.Redirect(w, r, "display", 301)
    }
}
