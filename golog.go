package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// configure logger
	f, err := os.OpenFile("urls.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", logURL)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func logURL(w http.ResponseWriter, r *http.Request) {
	// write URL to log file
	log.Println(r.Host + r.URL.Path)
	// write log file to response
	urls, err := ioutil.ReadFile("urls.log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", string(urls))
}
