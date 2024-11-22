package controllers

import (
	"PLP_Backend/classes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Estrutura para receber os nomes dos heróis na requisição
type SimulacaoBatalhaRequest struct {
	Heroi1 string `json:"heroi1"`
	Heroi2 string `json:"heroi2"`
}

// Controller para simular batalha entre heróis
func SimularBatalhaController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
		return
	}
	var req SimulacaoBatalhaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar requisição", http.StatusBadRequest)
		return
	}
	if req.Heroi1 == "" || req.Heroi2 == "" {
		http.Error(w, "Os nomes dos dois heróis são obrigatórios", http.StatusBadRequest)
		return
	}
	simulador := classes.SimuladorBatalha{}
	relatorio, err := simulador.SimularBatalhaComNomes(req.Heroi1, req.Heroi2)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro na simulação: %v", err), http.StatusInternalServerError)
		return
	}
	// Configura o cabeçalho e envia a resposta JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(relatorio)
}
