package main

import (
	"fmt"
	"math/rand"
)

// Heroi representa um herói com nome, força e popularidade
type Heroi struct {
	nome         string
	forca        int
	popularidade int
}

// Estrutura para armazenar o relatório de desempenho de um herói
type RelatorioHeroi struct {
	Nome           string `json:"nome"`
	ForcaInicial   int    `json:"forca_inicial"`
	Popularidade   int    `json:"popularidade"`
	ImpactoPopular int    `json:"impacto_popularidade"`
	FatorAleatorio int    `json:"fator_aleatorio"`
	AcertoCritico  bool   `json:"acerto_critico"`
	MoralNaLuta    int    `json:"moral_na_luta"`
	ForcaFinal     int    `json:"forca_final"`
}

// Estrutura para o relatório completo da batalha
type RelatorioBatalha struct {
	Vencedor string         `json:"vencedor"`
	Perdedor string         `json:"perdedor"`
	Heroi1   RelatorioHeroi `json:"heroi1"`
	Heroi2   RelatorioHeroi `json:"heroi2"`
}

// SimuladorBatalha simula uma batalha entre dois heróis
type SimuladorBatalha struct{}

// chanceDeAcertoCritico calcula a chance de acerto crítico com base na popularidade do herói
func (sb SimuladorBatalha) chanceDeAcertoCritico(popularidade int) float64 {
	// A chance de acerto crítico é proporcional à popularidade (máximo de 50%)
	chance := float64(popularidade) / 2.0
	if chance > 50 {
		chance = 50 // Limita a chance de crítico a 50%
	}
	return chance
}

// calculaMoral ajusta a moral com base na diferença de força durante a batalha
func (sb SimuladorBatalha) calculaMoral(forcaHeroi1, forcaHeroi2 int) int {
	// Se o herói 1 está ganhando, sua moral aumenta
	moral := 0
	if forcaHeroi1 > forcaHeroi2 {
		moral = 5
	} else if forcaHeroi2 > forcaHeroi1 {
		moral = -5
	}
	return moral
}

// BuscarHeroiParaBatalha busca um herói no banco de dados pelo nome
func BuscarHeroiParaBatalha(nomeHeroi string) (*Heroi, error) {
	db := ConectaDB()
	defer db.Close()

	query := `
		SELECT nome_heroi, forca, popularidade
		FROM herois
		WHERE nome_heroi = $1 AND esconder = false;
	`

	var heroi Heroi
	err := db.QueryRow(query, nomeHeroi).Scan(
		&heroi.nome,
		&heroi.forca,
		&heroi.popularidade,
	)

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar herói", err)
	}

	return &heroi, nil
}

// Modifica a função para retornar o relatório
func (sb SimuladorBatalha) simularBatalha(heroi1, heroi2 Heroi) RelatorioBatalha {
	// Calcula a força final de cada herói, considerando a popularidade
	forcaFinalHeroi1 := heroi1.forca + (heroi1.popularidade / 2)
	forcaFinalHeroi2 := heroi2.forca + (heroi2.popularidade / 2)

	// Verifica a chance de acerto crítico para cada herói
	chanceCriticoHeroi1 := sb.chanceDeAcertoCritico(heroi1.popularidade)
	chanceCriticoHeroi2 := sb.chanceDeAcertoCritico(heroi2.popularidade)

	// Variáveis para registrar se o herói teve um acerto crítico
	criticoHeroi1 := false
	criticoHeroi2 := false

	// Adiciona o impacto de acerto crítico
	if rand.Float64()*100 < chanceCriticoHeroi1 {
		criticoHeroi1 = true
		forcaFinalHeroi1 += 20 // Aumenta a força com um bônus de crítico
	}

	if rand.Float64()*100 < chanceCriticoHeroi2 {
		criticoHeroi2 = true
		forcaFinalHeroi2 += 20 // Aumenta a força com um bônus de crítico
	}

	// Inicializa fatores aleatórios
	fatorAleatorioHeroi1 := 0
	fatorAleatorioHeroi2 := 0

	// Adiciona um fator aleatório para cada herói (20% de chance)
	if rand.Intn(10) < 2 {
		fatorAleatorioHeroi1 = 15
		forcaFinalHeroi1 += fatorAleatorioHeroi1
	}
	if rand.Intn(10) < 2 {
		fatorAleatorioHeroi2 = 15
		forcaFinalHeroi2 += fatorAleatorioHeroi2
	}

	// Calcula a moral temporária com base nas forças relativas
	moralHeroi1 := sb.calculaMoral(forcaFinalHeroi1, forcaFinalHeroi2)
	moralHeroi2 := sb.calculaMoral(forcaFinalHeroi2, forcaFinalHeroi1)

	// Aplica a moral no cálculo final da força
	forcaFinalHeroi1 += moralHeroi1
	forcaFinalHeroi2 += moralHeroi2

	// Calcula a probabilidade de vitória do herói 1
	forcaTotal := forcaFinalHeroi1 + forcaFinalHeroi2
	chanceHeroi1 := float64(forcaFinalHeroi1) / float64(forcaTotal)

	// Determina o vencedor com base na probabilidade calculada
	vencedor := heroi2
	perdedor := heroi1
	if rand.Float64() < chanceHeroi1 {
		vencedor = heroi1
		perdedor = heroi2
	}
	fmt.Printf("Batalha concluida\n%s venceu a batalha", vencedor.nome)

	// Cria o relatório da batalha
	relatorio := RelatorioBatalha{
		Vencedor: vencedor.nome,
		Perdedor: perdedor.nome,
		Heroi1: RelatorioHeroi{
			Nome:           heroi1.nome,
			ForcaInicial:   heroi1.forca,
			Popularidade:   heroi1.popularidade,
			ImpactoPopular: heroi1.popularidade / 2,
			FatorAleatorio: fatorAleatorioHeroi1,
			AcertoCritico:  criticoHeroi1,
			MoralNaLuta:    moralHeroi1,
			ForcaFinal:     forcaFinalHeroi1,
		},
		Heroi2: RelatorioHeroi{
			Nome:           heroi2.nome,
			ForcaInicial:   heroi2.forca,
			Popularidade:   heroi2.popularidade,
			ImpactoPopular: heroi2.popularidade / 2,
			FatorAleatorio: fatorAleatorioHeroi2,
			AcertoCritico:  criticoHeroi2,
			MoralNaLuta:    moralHeroi2,
			ForcaFinal:     forcaFinalHeroi2,
		},
	}
	return relatorio
}

func (sb SimuladorBatalha) simularBatalhaComNomes(nomeHeroi1, nomeHeroi2 string) (*RelatorioBatalha, error) {
	heroi1, err := BuscarHeroiParaBatalha(nomeHeroi1)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar herói 1")
	}

	heroi2, err := BuscarHeroiParaBatalha(nomeHeroi2)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar herói 2")
	}

	relatorio := sb.simularBatalha(*heroi1, *heroi2)
	return &relatorio, nil
}
