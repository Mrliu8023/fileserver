package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed files
var efs embed.FS

func main() {
	port := os.Getenv("WEBPORT")
	if port == "" {
		port = ":8000"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	subfs, _ := fs.Sub(efs, "files")
	if err := http.ListenAndServe(port, http.FileServer(http.FS(subfs))); err != nil {
		log.Fatal(err)
	}
}
