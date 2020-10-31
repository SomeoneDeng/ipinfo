package main

import (
	"net/http"
	"strings"
)

func ipInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Split(r.RemoteAddr, ":")[0]))
}

func main() {
	http.HandleFunc("/ipinfo", ipInfo)
	http.ListenAndServe(":8765", nil)
}
