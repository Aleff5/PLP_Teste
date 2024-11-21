package controllers

import (
	"encoding/json"
	"net/http"
	"teste/classes"
)

func ConsultaMissaoHeroi(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		NomeHeroi string `json:"nome_heroi"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	nomeHeroi := requestData.NomeHeroi

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	missao, err := classes.ConsultaMissoesPorHeroi(nomeHeroi)
	if err != nil {
		http.Error(w, "Missão não encontrada ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(missao)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}
