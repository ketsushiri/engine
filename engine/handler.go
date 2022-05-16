package engine

import (
	"log"
	"net/http"
)

func Run(address string) {
	http.HandleFunc("/auth", AuthHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}
