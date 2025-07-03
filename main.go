package main

import (
	"github.com/iamharvey/weboo/internal"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", internal.HomeHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
