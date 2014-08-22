package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host + r.URL.Path)
	t := time.Now().Format("2006-1-2 3:04:05")
	fmt.Fprintf(w, "%s\t%s", t, r.Host+r.URL.Path)
}

func main() {
	// configure logger
	f, err := os.OpenFile("urls.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
