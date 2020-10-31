package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Info struct {
	Ip          string `json:"ip"`
	UpdatedTime int64  `json:"updated_time"`
}

func (info *Info) ipInfo(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(info)
	w.Write(data)
}

func (info *Info) report(w http.ResponseWriter, r *http.Request) {
	info.Ip = strings.Split(r.RemoteAddr, ":")[0]
	info.UpdatedTime = time.Now().Unix()
}

func main() {

	var info Info

	http.HandleFunc("/ipinfo", info.ipInfo)
	http.HandleFunc("/reportip", info.report)
	http.ListenAndServe(":8765", nil)
}
