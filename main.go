package main

import (
	"net/http"
	"os"
)

func main() {
	rootpath := os.Getenv("SERVERPATH")

	if err := http.ListenAndServe(":80", http.FileServer(http.Dir(rootpath))); err != nil {
		panic(err)
	}
}
