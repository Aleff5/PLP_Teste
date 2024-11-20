package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Loading() {
	r := mux.NewRouter()

	r.HandleFunc("/", MostraTudo)
	r.HandleFunc("/heroi", MostraPorNome)
	r.HandleFunc("/heroipop", MostraPopularidade)
	r.HandleFunc("/heroistatus", MostraPorStatus)
	log.Fatal(http.ListenAndServe(":8080", (r)))
}
