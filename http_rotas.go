package main

import (
	"log"
	"net/http"
	"teste/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Define as rotas da aplicação
func Loading() {
	r := mux.NewRouter()

	// Rotas para os herois
	// r.HandleFunc("/", controllers.MostraTudo).Methods("GET")
	r.HandleFunc("/", controllers.MostraTodosOsNomesHerois).Methods("GET")
	r.HandleFunc("/heroi", controllers.MostraPorNome).Methods("POST")
	r.HandleFunc("/heroipop", controllers.MostraPopularidade).Methods("POST")
	r.HandleFunc("/heroicadastra", controllers.CadastraHeroi).Methods("POST")

	r.HandleFunc("/delete", controllers.DeletaAKAralha).Methods("DELETE")

	r.HandleFunc("/heroistatus", controllers.MostraPorStatus).Methods("POST")
	r.HandleFunc("/poderes", controllers.MostraTodosPoderes).Methods("GET")
	r.HandleFunc("/editar", controllers.EditarHeroiHandler).Methods("POST")
	// Rotas para os crimes
	r.HandleFunc("/heroieseveridadecrime", controllers.ConsultaCrimesHS).Methods("POST")
	r.HandleFunc("/heroicrime", controllers.ConsultaCrimesHeroi).Methods("POST")
	r.HandleFunc("/severidadecrime", controllers.ConsultaCrimesSeveridade).Methods("POST")
	// Rotas para as missoes
	r.HandleFunc("/missao", controllers.ConsultaMissaoHeroi).Methods("POST")

	// suporte a CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins, credentials)(r)))

	log.Fatal(http.ListenAndServe(":8080", (r)))
}
