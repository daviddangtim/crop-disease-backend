package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
log.Println("File Upload Endpoint Hit")

r.ParseMultipartForm(10 << 20)

file, handler, err := r.FormFile("image")
if err != nil {
    log.Println(http.StatusBadRequest)
    log.Println(err)
    return
}
defer file.Close()
log.Printf("Uploaded File: %+v\n",handler.Filename)
log.Printf("File Size: %+v\n",handler.Size)
log.Printf("MIME Header: %+v\n",handler.Header)

tempFile, err := os.CreateTemp("temp-images", "upload-*.png")
if err != nil {
    log.Println(err)
}
defer tempFile.Close()

fileBytes, err := io.ReadAll(file)
if err != nil {
    log.Println(err)
}

tempFile.Write(fileBytes)
fmt.Fprintf(w, "Successfully Uploaded File\n")

    
}



func startServer()  {
    mux := http.NewServeMux()
    mux.HandleFunc("/upload", uploadFile)
    s := &http.Server{
        Addr: ":8080",
        Handler: mux,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20, //Header size of 1mb
    }
    log.Fatal(s.ListenAndServe())    
}

func main() {
   startServer()
}