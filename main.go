package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ResponseInfo struct {
	URL    string              `json:"url"`
	Path   string              `json:"path"`
	Query  string              `json:"query"`
	Method string              `json:"method"`
	Body   []byte              `json:"body"`
	Header map[string][]string `json:"header"`
}

func handleAny(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp ResponseInfo
	resp.Method = r.Method
	resp.URL = r.URL.String()
	resp.Path = r.URL.Path
	resp.Query = r.URL.RawQuery

	// body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	resp.Body = body

	// header
	header := make(map[string][]string)
	for key, value := range r.Header {
		header[key] = value
	}
	resp.Header = header

	json.NewEncoder(w).Encode(resp)
}

func main() {
	handler := http.HandlerFunc(handleAny)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
