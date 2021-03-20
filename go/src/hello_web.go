package main

import (
	"io"
	"log"
	"net/http"
	"time"
)
func Greeting(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "<b>Code.education Rocks!!</b>\n")
}
func Hello(w http.ResponseWriter, req *http.Request)  {
		io.WriteString(w, "Hello!\n")
}
func main() {
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/", Greeting)
	http.HandleFunc("/healthcheck", func(rw http.ResponseWriter, r *http.Request) { io.WriteString(rw, "I'm healthy") })
	log.Fatal(s.ListenAndServe())
}