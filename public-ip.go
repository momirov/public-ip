package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handlerText)
	http.HandleFunc("/json", handlerJson)
	http.ListenAndServe(":8083", nil)
}

func handlerText(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr,":")[0]
	fmt.Fprintf(w, ip)
}

func handlerJson(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr,":")[0]
	fmt.Fprintf(w, "{ \"ip\": \"%s\" }" ,ip)
}
