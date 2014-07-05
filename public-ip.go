package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"
)

func main() {
	http.HandleFunc("/", handlerText)
	http.HandleFunc("/json", handlerJson)
	http.ListenAndServe(GetPort(), nil)
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8083"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func GetRemoteAddr(r *http.Request) string {
	ip := strings.Split(r.RemoteAddr,":")[0]

	return ip
}

func handlerText(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, GetRemoteAddr(r))
}

func handlerJson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{ \"ip\": \"%s\" }", GetRemoteAddr(r))
}
