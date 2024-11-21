package main

import (
	"encoding/json"
	"fmt"

	//"fmt"
	"net/http"
)

// Controller para exibir todas as informações de todos os herois
func MostraTudo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var herois Herois
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

	herois, err := BuscaHeroisPorPopularidade(popularidade)
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

	herois, err := BuscaHeroisPorStatus(status)
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
		Heroi      Herois `json:"heroi"`
		IDsPoderes []int  `json:"ids_poderes"` // Agora recebemos apenas os IDs dos poderes
	}

	// Decodifica o JSON da requisição
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Payload da requisição inválido", http.StatusBadRequest)
		return
	}

	// Chama a função para cadastrar o herói com os IDs dos poderes
	err = CadastrarHeroiComPoderesNormalizados(requestPayload.Heroi, requestPayload.IDsPoderes)
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
	possivelerro := Remove(id)
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

// Controller para consultar crimes por heroi e severidade
func ConsultaCrimesHS(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		NomeHeroi        string `json:"nome_heroi"`
		SeveridadeMinima int    `json:"severidade_minima"`
		SeveridadeMaxima int    `json:"severidade_maxima"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	nomeHeroi := requestData.NomeHeroi
	severidadeMinima := requestData.SeveridadeMinima
	severidadeMaxima := requestData.SeveridadeMaxima

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	crimes, err := ConsultaCrimesPorHeroiESeveridade(nomeHeroi, severidadeMinima, severidadeMaxima)
	if err != nil {
		http.Error(w, "Crimes não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(crimes)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

// Controller para consultar crimes por heroi
func ConsultaCrimesHeroi(w http.ResponseWriter, r *http.Request) {
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

	crimes, err := ConsultaCrimesPorHeroi(nomeHeroi)
	if err != nil {
		http.Error(w, "Crimes não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(crimes)
	if err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

// Controller para consultar todos os poderes e seus IDs
func MostraTodosPoderes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allPoderes := ExibeTodosOsPoderes()
	json.NewEncoder(w).Encode(allPoderes)
}

// Controller para consultar crimes de acordo com a severidade
func ConsultaCrimesSeveridade(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		SeveridadeMinima int `json:"severidade_minima"`
		SeveridadeMaxima int `json:"severidade_maxima"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	severidadeMinima := requestData.SeveridadeMinima
	severidadeMaxima := requestData.SeveridadeMaxima

	// Configura o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	crimes, err := ConsultaCrimesPorSeveridade(severidadeMinima, severidadeMaxima)
	if err != nil {
		http.Error(w, "Crimes não encontrado ou erro no servidor", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(crimes)
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
		NomeHeroi       string `json:"nome_heroi"`       // Nome do herói a ser editado
		HeroiAtualizado Herois `json:"heroi_atualizado"` // Dados atualizados do herói
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
	err = EditarHeroiPorNome(requestPayload.NomeHeroi, requestPayload.HeroiAtualizado)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao editar herói: %v", err), http.StatusInternalServerError)
		return
	}

	// Retorna uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Herói atualizado com sucesso!"))
}

// Controller para consultar missões por heroi
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

	missao, err := ConsultaMissoesPorHeroi(nomeHeroi)
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
