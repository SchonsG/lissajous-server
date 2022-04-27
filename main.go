// Server echo the HTTP request
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		cycle := r.FormValue("cycle")

		if s, err := strconv.ParseFloat(cycle, 64); err == nil {
			lissajous(w, s)
		}

	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echo the path component from the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
