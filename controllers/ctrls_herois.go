package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"teste/classes"
)

// Controller para exibir todas as informações de todos os herois
func MostraTudo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var herois classes.Herois
	allHeroes := herois.ExibeInfosGerais()
	json.NewEncoder(w).Encode(allHeroes)

}

// Controller para exibir todas as informações de um heroi por nome
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

	heroi, err := classes.BuscaHeroiPorNome(nomeHeroi)
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

// Controller para exibir todas as informações de um heroi por Popularidade
func MostraPopularidade(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Popularidade int `json:"popularidade"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	popularidade := requestData.Popularidade

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	herois, err := classes.BuscaHeroisPorPopularidade(popularidade)
	if err != nil {
		http.Error(w, "Herois não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(herois)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

// Controller para exibir todas as informações de um heroi por Status de atividade
func MostraPorStatus(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Status string `json:"status_atividade"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	status := requestData.Status

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	herois, err := classes.BuscaHeroisPorStatus(status)
	if err != nil {
		http.Error(w, "Herois não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(herois)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

// Controller para cadastrar um heroi
func CadastraHeroi(w http.ResponseWriter, r *http.Request) {
	// Estrutura para decodificar o payload
	var requestPayload struct {
		Heroi      classes.Herois `json:"heroi"`
		IDsPoderes []int          `json:"ids_poderes"` // Agora recebemos apenas os IDs dos poderes
	}

	// Decodifica o JSON da requisição
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Payload da requisição inválido", http.StatusBadRequest)
		return
	}

	// Chama a função para cadastrar o herói com os IDs dos poderes
	err = classes.CadastrarHeroiComPoderesNormalizados(requestPayload.Heroi, requestPayload.IDsPoderes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao cadastrar herói: %v", err), http.StatusInternalServerError)
		return
	}

	// Resposta de sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Herói cadastrado com sucesso!"))
}

// Controller para deletar um heroi
func DeletaAKAralha(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Id int `json:"id_heroi"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	id := requestData.Id
	possivelerro := classes.Remove(id)
	if possivelerro != nil {
		http.Error(w, "Herois não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode("tudo OK")
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}

}

// Handler para editar um heroi
func EditarHeroiHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método da requisição é PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido. Use PUT.", http.StatusMethodNotAllowed)
		return
	}

	// Estrutura para decodificar o payload da requisição
	var requestPayload struct {
		NomeHeroi       string         `json:"nome_heroi"`       // Nome do herói a ser editado
		HeroiAtualizado classes.Herois `json:"heroi_atualizado"` // Dados atualizados do herói
	}

	// Decodifica o JSON do corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Payload da requisição inválido", http.StatusBadRequest)
		return
	}

	// Verifica se o nome do herói foi fornecido
	if requestPayload.NomeHeroi == "" {
		http.Error(w, "O nome do herói deve ser fornecido", http.StatusBadRequest)
		return
	}

	// Chama a função para editar os dados do herói
	err = classes.EditarHeroiPorNome(requestPayload.NomeHeroi, requestPayload.HeroiAtualizado)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao editar herói: %v", err), http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Herói atualizado com sucesso!"))
}

// Controller para consultar todos os poderes e seus IDs
func MostraTodosPoderes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allPoderes := classes.ExibeTodosOsPoderes()
	json.NewEncoder(w).Encode(allPoderes)
}
