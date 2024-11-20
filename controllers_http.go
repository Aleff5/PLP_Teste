package main

import (
	"encoding/json"
	"net/http"
)

func MostraTudo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var herois Herois
	allHeroes := herois.ExibeInfosGerais()
	json.NewEncoder(w).Encode(allHeroes)

}

func MostraPorNome(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		NomeHeroi string `json:"nome_heroi"`
	}

	// Decodifica o JSON do corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	nomeHeroi := requestData.NomeHeroi

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	heroi, err := BuscaHeroiPorNome(nomeHeroi)
	if err != nil {
		http.Error(w, "Herói não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}

	// Codifica e envia a resposta JSON
	err = json.NewEncoder(w).Encode(heroi)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}
