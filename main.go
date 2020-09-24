package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/", ExampleHandler)
	http.HandleFunc("/test-push", ExampleHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	LogRequest(r)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}

func LogRequest(r *http.Request) {
	log.Println(formatRequest(r))
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string

	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%s: %s", name, h))
		}
	}

	buf, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println("bodyErr ", bodyErr.Error())

		return strings.Join(request, "\n")
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	request = append(request, fmt.Sprintf("BODY: %v\n---------------\n", rdr1))

	return strings.Join(request, "\n")
}