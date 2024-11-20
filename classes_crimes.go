package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Crimes struct {
	Herois
	NomeCrime       string `json:"nome_crime"`
	Severidade      string `json:"severidade"`
	DataCrime       string `json:"data_crime"`
	DescricaoEvento string `json:"descricao_evento"`
}

type Missoes struct {
	NomeMissao      string `json:"nome_missao"`
	DescricaoMissao string `json:"descricao"`
	NivelMissao     string `json:"nivel_dificuldade"`
	Resultado       string `json:"resultado"`
	Recompensa      string `json:"recompensa"`
}

// Método para consultar crimes por herói e por severidade
func ConsultaCrimesPorHeroiESeveridade(nomeHeroi string, severidadeMinima int, severidadeMaxima int) ([]Crimes, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base no nome do herói e na severidade
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento, hc.esconder
		FROM 
			Crimes c
		JOIN 
			Herois_Crimes hc ON c.id_crime = hc.id_crime
		JOIN 
			Herois h ON hc.id_heroi = h.id_heroi
		WHERE 
			h.nome_heroi = $1 
		AND 
			hc.esconder = false
		AND 
			c.severidade BETWEEN $2 AND $3;
	`

	// Executa a consulta
	rows, err := db.Query(query, nomeHeroi, severidadeMinima, severidadeMaxima)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado para o herói %s com severidade entre %d e %d", nomeHeroi, severidadeMinima, severidadeMaxima)
	}

	return crimes, nil
}

// Função para Consultar os Crimes por Herói
func ConsultaCrimesPorHeroi(nomeHeroi string) ([]Crimes, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base no nome do herói
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento
		FROM 
			Crimes c
		JOIN 
			Herois_Crimes hc ON c.id_crime = hc.id_crime
		JOIN 
			Herois h ON hc.id_heroi = h.id_heroi
		WHERE 
			h.nome_heroi ILIKE $1
		AND 
			hc.esconder = false;
	`

	// Executa a consulta
	rows, err := db.Query(query, nomeHeroi)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado para o herói %s", nomeHeroi)
	}

	return crimes, nil
}

// Função para Consultar os Crimes por Severidade
func ConsultaCrimesPorSeveridade(severidadeMinima int, severidadeMaxima int) ([]Crimes, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base na severidade
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento
		FROM 
			Crimes c
		JOIN 
			Herois_Crimes hc ON c.id_crime = hc.id_crime
		WHERE 
			c.severidade BETWEEN $1 AND $2
		AND 
			hc.esconder = false;
	`

	// Executa a consulta
	rows, err := db.Query(query, severidadeMinima, severidadeMaxima)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado com severidade entre %d e %d", severidadeMinima, severidadeMaxima)
	}

	return crimes, nil
}

// Função para Modificar Informações do Herói
func ModificacaoHeroi(NomeHeroi string, NovoNomeHeroi string, NomeReal string, Sexo string, Altura float64, Local_nascimento string, Data_nascimento float64, Peso float64, Popularidade int, Forca int, Status string) error {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso
	// Query que atualiza o herói com "NomeHeroi"
	query := `
		UPDATE
			Herois
		SET
			nome_heroi = $2,
			nome = $3,
			sexo = $4,
			altura = $5,
			local_nasc = $6,
			data_nasc = $7,
			peso = $8,
			popularidade = $9,
			forca = $10,
			status = $11;
		WHERE
			nome_heroi = $1
			`
	rows, err := db.Query(query, NomeHeroi, NovoNomeHeroi, NomeReal, Sexo, Altura, Local_nascimento, Data_nascimento, Peso, Popularidade, Forca, Status)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso
	return nil
}

// Função para Consultar Missões por Herói
func ConsultaMissoesPorHeroi(nomeHeroi string) ([]Missoes, error) {
	db := ConectaDB()
	defer db.Close() //Grarantir que o banco de dados seja fechado após o uso
	//Query para buscar missões com base no nome do herói
	query := `
		SELECT
			m.nome_missao, m.descricao, m.nivel_dificuldade, m.resultado, m.recompensa
		FROM
			Missoes m
		JOIN
			Herois_Missoes hm ON m.id_missao = hm.id_missao
		JOIN
			Herois h ON hm.id_heroi = h.id_heroi
		WHERE
			h.nome_heroi = $1;
		ORDER BY m.nivel_dificuldade ASC;
	`
	//Executa a consulta
	rows, err := db.Query(query, nomeHeroi)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()    //Garantir que o resultado seja fechado após o uso
	var missoes []Missoes //Cria uma slice para armazenar as missões
	//Itera sobre os resultados da consulta
	for rows.Next() {
		var missao Missoes
		err := rows.Scan(
			&missao.NomeMissao,
			&missao.DescricaoMissao,
			&missao.NivelMissao,
			&missao.Resultado,
			&missao.Recompensa,
		)
		if err != nil {
			log.Fatal(err)
		}
		missoes = append(missoes, missao)
	}
	//Verifica se não encontrou nenhuma missão
	if len(missoes) == 0 {
		return nil, fmt.Errorf("nenhuma missão encontrada para o herói %s", nomeHeroi)
	}
	return missoes, nil
}
