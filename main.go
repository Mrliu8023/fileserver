package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("WEBPORT")
	if port == "" {
		port = ":8000"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	if err := http.ListenAndServe(port, http.FileServer(http.Dir("files"))); err != nil {
		log.Fatal(err)
	}
}
