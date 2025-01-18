package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handleAny(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%e", err)
		return
	}
	fmt.Fprintf(w, "%s", dump)
}

func main() {
	handler := http.HandlerFunc(handleAny)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
