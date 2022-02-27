package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	bindhost = flag.String("bind", ":3031", "host/port to listen on")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from nix!")
	})

	log.Printf("listening for HTTP on %s", *bindhost)
	log.Fatal(http.ListenAndServe(*bindhost, nil))
}
