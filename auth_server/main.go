package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)




func startServer()  {
    mux := http.NewServeMux()
    // mux.HandleFunc("/", )
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