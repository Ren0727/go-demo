package main

import (
    "fmt"
    "log"
    "net/http"
)

const version = "v1.0.0"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        fmt.Fprintf(w, "Go 镜像部署练习，当前版本：%s\n", version)
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        fmt.Fprintln(w, "ok")
    })

    log.Printf("启动 go-demo，版本 %s，监听端口 8080", version)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
