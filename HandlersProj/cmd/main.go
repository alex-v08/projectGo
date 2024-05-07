package main

import (
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + t))
}

func main() {

	// Código principal aquí

	mux := http.NewServeMux()
	th := timeHandler{format: time.RFC1123}

	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)

}
