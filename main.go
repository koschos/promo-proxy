package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rssHandler)
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
