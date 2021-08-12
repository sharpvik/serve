package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var addr string
	port := flag.Int("port", 8080, "Address to serve at")
	pub := flag.Bool("pub", false, "Serve on public interface")
	flag.Parse()

	if *pub {
		addr = fmt.Sprintf("0.0.0.0:%d", *port)
	} else {
		addr = fmt.Sprintf("localhost:%d", *port)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(cwd))))
	log.Printf("serving at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
