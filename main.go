package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const HelpText = `Usage: whttp <host> [<port>]

Provide the port number as part of the host or as a single argument.
For example: localhost:5000, :5000, or localhost 5000.`

// This function simply prints a dump of the http request.
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
	var addr string

	if len(os.Args[1:]) == 1 {
		addr = os.Args[1]
	} else if len(os.Args[1:]) == 2 {
		addr = fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])
	} else {
		log.Fatal(HelpText)
		return
	}

	handler := http.HandlerFunc(handleAny)
	log.Fatal(http.ListenAndServe(addr, handler))
}
