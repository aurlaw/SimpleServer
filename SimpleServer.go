package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var chttp = http.NewServeMux()

func main() {

	homeDir := flag.String("dir", "./", "directory of files to serve")
	port := flag.Int("port", 8080, "port to listen on")

	flag.Parse()

	fmt.Println("dir:", *homeDir)
	fmt.Println("port:", *port)

	chttp.Handle("/", http.FileServer(http.Dir(*homeDir)))

	http.HandleFunc("/", HomeHandler) // homepage
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, ".") {
		chttp.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(w, "HomeHandler")
	}
}
