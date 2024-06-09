package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	// Try to serve from ./static/pages first
	pagesPath := filepath.Join("./static/pages", r.URL.Path)

	if fileInfo, err := os.Stat(pagesPath); err == nil && fileInfo.IsDir() {
		pagesPath = filepath.Join(pagesPath, "index.html")
		http.ServeFile(w, r, pagesPath)
		return
	}

	// serve from ./static
	staticPath := filepath.Join("./static", r.URL.Path)
	if _, err := os.Stat(staticPath); err == nil {
		http.ServeFile(w, r, staticPath)
		return
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", staticHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
