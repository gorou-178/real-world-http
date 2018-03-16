package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hellow</dody></html>")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listening ::18443")
	err := http.ListenAndServeTLS("example.com:18443", "../../certificate/server.crt", "../../certificate/server.key", nil)
	log.Println(err)
}
