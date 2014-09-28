// simple filserver written in go for demo purposes
package main

import (
	"flag"
	"net/http"
)

var port = flag.String("port", "9090", "port to bind to")

func main() {
	flag.Parse()
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("."))))
}