package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	host string
	mode string
)

func init() {
	log.SetPrefix("[ Cmode ]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	var err error
	host, err = os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	mode = os.Getenv("MODE")
	if mode == "" {
		mode = "DEFAULT"
	}
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	a := v.Get("foo")
	fmt.Println("Request hit", a)
	fmt.Fprintf(w, "Welcome you , %s ! From  %s , Mode %s", a, host, mode)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ding")
}

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
