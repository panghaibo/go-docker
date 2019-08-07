package main

import (
        "fmt"
        "net"
        "net/http"
        "time"
)

func main() {
        serverMux := http.NewServeMux()
        serverMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello,world")
        })

        serverMux.HandleFunc("/pms/address", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "PMS,ADDRESS")
        })

        addr,_ := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
        listener, _ := net.ListenTCP("tcp", addr)

        s := &http.Server{
                Addr:           ":8080",
                Handler:        serverMux,
                ReadTimeout:    10 * time.Second,
                WriteTimeout:   10 * time.Second,
                MaxHeaderBytes: 1 << 20,
        }
        s.Serve(listener)
        //log.Fatal(s.ListenAndServe())
}
