package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-1-2 3:04:05")
	fmt.Fprintf(w, "%s   %s", t, r.Host+r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
