package classes

import (
	"fmt"
	"log"
	"teste/database"
)

// Struct de Missoes para o rows
type Missoes struct {
	NomeHeroi       string `json:"nome_heroi"`
	NomeMissao      string `json:"nome_missao"`
	DescricaoMissao string `json:"descricao"`
	NivelMissao     string `json:"nivel_dificuldade"`
	Resultado       string `json:"resultado"`
	Recompensa      string `json:"recompensa"`
}

// Função para Consultar Missões por Herói
func ConsultaMissoesPorHeroi(nomeHeroi string) ([]Missoes, error) {
	db := database.ConectaDB()
	defer db.Close()

	// Query atualizada para incluir todos os heróis da missão
	query := `
		SELECT
			m.nome_missao, 
			m.descricao, 
			m.nivel_dificuldade, 
			m.resultado, 
			m.recompensa, 
			h.nome_heroi
		FROM
			missoes m
		JOIN
			herois_missoes hm ON m.id_missao = hm.id_missao
		JOIN
			herois h ON hm.id_heroi = h.id_heroi
		WHERE
			m.id_missao IN (
				SELECT DISTINCT hm.id_missao 
				FROM herois_missoes hm
				JOIN herois h ON hm.id_heroi = h.id_heroi
				WHERE h.nome_heroi = $1
			)
		ORDER BY m.nivel_dificuldade ASC;
	`

	rows, err := db.Query(query, nomeHeroi)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	//Itera sobre o resultado das consultas
	var missoes []Missoes
	for rows.Next() {
		var missao Missoes
		err := rows.Scan(
			&missao.NomeMissao,
			&missao.DescricaoMissao,
			&missao.NivelMissao,
			&missao.Resultado,
			&missao.Recompensa,
			&missao.NomeHeroi,
		)
		if err != nil {
			log.Fatal(err)
		}
		missoes = append(missoes, missao)
	}

	if len(missoes) == 0 {
		return nil, fmt.Errorf("nenhuma missão encontrada para o herói %s", nomeHeroi)
	}
	return missoes, nil
}
