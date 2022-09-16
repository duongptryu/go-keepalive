package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive a request from:", r.RemoteAddr, r.Header)
	w.Write([]byte("ok"))
}

func main() {
	var s = http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(Index),
	}
	fmt.Println("Server listen in port 8080")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}