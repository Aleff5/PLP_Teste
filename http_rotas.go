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
	r.HandleFunc("/heroicadastra", CadastraHeroi)
	r.HandleFunc("/delete", DeletaAKAralha)
	r.HandleFunc("/heroistatus", MostraPorStatus)
	r.HandleFunc("/heroieseveridadecrime", ConsultaCrimesHS)
	r.HandleFunc("/heroicrime", ConsultaCrimesHeroi)
	r.HandleFunc("/poderes", MostraTodosPoderes)
	r.HandleFunc("/severidadecrime", ConsultaCrimesSeveridade)
	r.HandleFunc("/editar", EditarHeroiHandler)
	r.HandleFunc("/missao", ConsultaMissaoHeroi)
	log.Fatal(http.ListenAndServe(":8080", (r)))
}
