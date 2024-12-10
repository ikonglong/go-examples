package main

import (
	"log"
	"net/http"

	"github.com/ikonglong/go-examples/concur/patterns/ctx/server"
)

func main() {
	http.HandleFunc("/search", server.HandleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
