package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.Int("addr", 8080, "address to serve at")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(cwd))))
	log.Printf("serving at localhost:%d", *addr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *addr), nil))
}
