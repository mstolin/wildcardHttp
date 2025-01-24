package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const HelpText = `Usage: whttp <host> [<JSON>]

host : The TCP network host to accept connections on; e.g. localhost:5000 or :5000
JSON : Serve any response in JSON format; Default: false

For example: wget :5000 true`

const PrettyHeaderFormat = "Header\n======\n%s"
const PrettyBodyFormat = "Body\n====\n%s"

type RequestData struct {
	Header map[string][]string
	Body   string
}

func (data RequestData) prettyPrintHeader() string {
	var output string
	for key, values := range data.Header {
		output += fmt.Sprintf("%s: %s\n", key, strings.Join(values, ", "))
	}
	return output
}

func (data RequestData) PrettyPrint() string {
	prettyHeader := fmt.Sprintf(PrettyHeaderFormat, data.prettyPrintHeader())
	var prettyBody string
	if len(data.Body) > 0 {
		prettyBody = fmt.Sprintf(PrettyBodyFormat, data.Body)
	}
	return fmt.Sprintf("%s\n%s", prettyHeader, prettyBody)
}

func NewRequestData(r *http.Request) (RequestData, error) {
	var data RequestData
	data.Header = r.Header
	if body, err := io.ReadAll(r.Body); err == nil {
		data.Body = string(body)
	} else {
		return data, err
	}
	return data, nil
}

// This function simply prints a dump of the http request.
func handleAny(w http.ResponseWriter, r *http.Request) {
	if requestData, err := NewRequestData(r); err == nil {
		fmt.Fprintf(w, "%s", requestData.PrettyPrint())
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%e", err)
		return
	}
}

func handleAnyJson(w http.ResponseWriter, r *http.Request) {
	if requestData, err := NewRequestData(r); err == nil {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(requestData); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "%e", err)
			return
		}
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%e", err)
		return
	}
}

func main() {
	var addr string
	var serveJson bool

	if len(os.Args[1:]) == 1 {
		addr = os.Args[1]
	} else if len(os.Args[1:]) == 2 {
		addr = os.Args[1]
		if _serveJson, err := strconv.ParseBool(os.Args[2]); err == nil {
			serveJson = _serveJson
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(HelpText)
	}

	var handler http.Handler
	if serveJson {
		handler = http.HandlerFunc(handleAnyJson)
	} else {
		handler = http.HandlerFunc(handleAny)
	}
	log.Fatal(http.ListenAndServe(addr, handler))
}
