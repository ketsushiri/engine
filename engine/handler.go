package engine

import (
	"log"
	"net/http"
)

func Run(address string) {
	http.HandleFunc("/auth", AuthHandler)
	http.HandleFunc("/register", RegisterHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}
