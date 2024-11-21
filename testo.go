package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// Função para conectar ao banco de dados
func ConectaDB() *sql.DB {
	conexao := "user=postgres dbname=TheBoys password=davi252310 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Estrutura de informações pessoais
type InfosPessoas struct {
	Nome      string    `json:"nome_real"`
	Sexo      string    `json:"sexo"`
	Peso      float64   `json:"peso"`
	Altura    float64   `json:"altura"`
	DataNasc  time.Time `json:"data_nascimento"`
	LocalNasc string    `json:"local_nascimento"`
}

// Estrutura dos Heróis
type Herois struct {
	InfosPessoas
	NomeHeroi    string   `json:"nome_heroi"`
	Poderes      []string `json:"poder"`
	Popularidade int      `json:"popularidade"`
	Status       string   `json:"status_atividade"`
	Forca        int      `json:"forca"`
}

// Estrutura dos Poderes
type Poder struct {
	Id_poder  int    `json:"id_poder"`
	Poder     string `json:"poder"`
	Descricao string `json:"descricao"`
}

// Método para exibir as informações dos heróis
func (h Herois) ExibeInfosGerais() []Herois {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	query := `
		SELECT 
			h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca, 
			COALESCE(STRING_AGG(p.poder, ', '), '') AS poderes
		FROM 
			Herois h
		LEFT JOIN 
			Herois_Poderes hp ON h.id_heroi = hp.id_heroi
		LEFT JOIN
			Poderes p ON p.id_poder = hp.id_poder
		WHERE 
			h.hide = false
		GROUP BY 
			h.id_heroi, h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca;
	`

	// Executa a consulta
	allInfos, err := db.Query(query)
	if err != nil {
		log.Fatalf("Erro ao executar a consulta: %v", err)
	}
	defer allInfos.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os heróis
	var informacoes []Herois

	// Itera sobre os resultados da consulta
	for allInfos.Next() {
		var heroi Herois
		var poderes *string     // Use ponteiro para lidar com valores NULL
		var dataNasc *time.Time // Use ponteiro para tratar data_nascimento como NULL

		err := allInfos.Scan(
			&heroi.Nome,
			&heroi.Sexo,
			&heroi.Peso,
			&heroi.Altura,
			&dataNasc, // Data de nascimento como ponteiro
			&heroi.LocalNasc,
			&heroi.NomeHeroi,
			&heroi.Popularidade,
			&heroi.Status,
			&heroi.Forca,
			&poderes,
		)
		if err != nil {
			log.Fatalf("Erro ao fazer o scan dos resultados: %v", err)
		}

		// Verifica se a data de nascimento é NULL
		if dataNasc != nil {
			heroi.DataNasc = *dataNasc // Converte para o valor de time.Time
		} else {
			heroi.DataNasc = time.Time{} // Define um valor padrão, se necessário
		}

		// Verifica se poderes é NULL e ajusta para uma string vazia se necessário
		if poderes != nil {
			heroi.Poderes = splitPoderes(*poderes)
		} else {
			heroi.Poderes = []string{} // Nenhum poder registrado
		}

		// Adiciona o herói à lista
		informacoes = append(informacoes, heroi)
	}

	// Verifica se ocorreu algum erro durante a iteração
	if err = allInfos.Err(); err != nil {
		log.Fatalf("Erro durante a iteração dos resultados: %v", err)
	}

	return informacoes
}

// Função para dividir poderes em uma slice
func splitPoderes(poderes string) []string {
	if poderes == "" {
		return []string{}
	}
	return strings.Split(poderes, ", ")
}

// Método para exibir as informações dos heróis por nome
func BuscaHeroiPorNome(nomeHeroi string) (*Herois, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar um herói específico pelo nome do herói
	query := `
		SELECT 
			h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca, 
			STRING_AGG(p.poder, ', ') AS poderes
		FROM 
			Herois h
		LEFT JOIN 
			Herois_Poderes hp ON h.id_heroi = hp.id_heroi
		LEFT JOIN
			Poderes p ON p.id_poder = hp.id_poder
		WHERE 
			h.nome_heroi = $1
		AND
			h.hide = false
		GROUP BY 
			h.id_heroi, h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca;
	`

	// Executa a consulta
	var heroi Herois
	var poderes *string     // Poderes como string, pode ser NULL
	var dataNasc *time.Time // Data de nascimento, pode ser NULL

	err := db.QueryRow(query, nomeHeroi).Scan(
		&heroi.Nome,
		&heroi.Sexo,
		&heroi.Peso,
		&heroi.Altura,
		&dataNasc,
		&heroi.LocalNasc,
		&heroi.NomeHeroi,
		&heroi.Popularidade,
		&heroi.Status,
		&heroi.Forca,
		&poderes,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("herói com nome %s não encontrado", nomeHeroi)
		}
		return nil, err
	}

	// Converte a data de nascimento, se não for NULL
	if dataNasc != nil {
		heroi.DataNasc = *dataNasc
	} else {
		heroi.DataNasc = time.Time{} // Define como valor zero
	}

	// Converte os poderes em uma slice, se não for NULL
	if poderes != nil {
		heroi.Poderes = splitPoderes(*poderes)
	} else {
		heroi.Poderes = []string{} // Nenhum poder registrado
	}

	return &heroi, nil
}

// Método para exibir as informações dos heróis por popularidade
func BuscaHeroisPorPopularidade(popularidade int) ([]Herois, error) {
	db := ConectaDB()
	defer db.Close()

	query := `
		SELECT 
			h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca, 
			STRING_AGG(p.poder, ', ') AS poderes
		FROM 
			Herois h
		LEFT JOIN 
			Herois_Poderes hp ON h.id_heroi = hp.id_heroi
		LEFT JOIN
			Poderes p ON p.id_poder = hp.id_poder
		WHERE 
			h.popularidade <= $1
		AND
			h.hide = false
		GROUP BY 
			h.id_heroi, h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca;
	`

	rows, err := db.Query(query, popularidade)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var herois []Herois
	for rows.Next() {
		var heroi Herois
		var poderes *string // Poderes como string, pode ser NULL
		var dataNasc *time.Time
		err := rows.Scan(
			&heroi.Nome,
			&heroi.Sexo,
			&heroi.Peso,
			&heroi.Altura,
			&dataNasc,
			&heroi.LocalNasc,
			&heroi.NomeHeroi,
			&heroi.Popularidade,
			&heroi.Status,
			&heroi.Forca,
			&poderes,
		)
		if err != nil {
			return nil, err
		}
		// Converte a data de nascimento, se não for NULL
		if dataNasc != nil {
			heroi.DataNasc = *dataNasc
		} else {
			heroi.DataNasc = time.Time{} // Define como valor zero
		}

		// Converte os poderes em uma slice, se não for NULL
		if poderes != nil {
			heroi.Poderes = splitPoderes(*poderes)
		} else {
			heroi.Poderes = []string{} // Nenhum poder registrado
		}

		herois = append(herois, heroi)
	}

	return herois, nil
}

// Método para exibir as informações dos heróis por status
func BuscaHeroisPorStatus(status string) ([]Herois, error) {
	db := ConectaDB()
	defer db.Close()

	// Consulta SQL para buscar heróis pelo status
	query := `
		SELECT 
			h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca, 
			STRING_AGG(p.poder, ', ') AS poderes
		FROM 
			Herois h
		LEFT JOIN 
			Herois_Poderes hp ON h.id_heroi = hp.id_heroi
		LEFT JOIN
			Poderes p ON p.id_poder = hp.id_poder
		WHERE 
			h.status_atividade = $1
		AND
			h.hide = false
		GROUP BY 
			h.id_heroi, h.nome_real, h.sexo, h.peso, h.altura, h.data_nascimento, h.local_nascimento, 
			h.nome_heroi, h.popularidade, h.status_atividade, h.forca;
	`

	rows, err := db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var herois []Herois
	for rows.Next() {
		var dataNasc *time.Time
		var poderes *string
		var heroi Herois
		err := rows.Scan(
			&heroi.Nome,
			&heroi.Sexo,
			&heroi.Peso,
			&heroi.Altura,
			&dataNasc,
			&heroi.LocalNasc,
			&heroi.NomeHeroi,
			&heroi.Popularidade,
			&heroi.Status,
			&heroi.Forca,
			&poderes,
		)
		if err != nil {
			return nil, err
		}
		// Converte a data de nascimento, se não for NULL
		if dataNasc != nil {
			heroi.DataNasc = *dataNasc
		} else {
			heroi.DataNasc = time.Time{} // Define como valor zero
		}

		//Converte os poderes em uma slice, se não for NULL
		if poderes != nil {
			heroi.Poderes = splitPoderes(*poderes)
		} else {
			heroi.Poderes = []string{} // Nenhum poder registrado
		}
		herois = append(herois, heroi)
	}

	return herois, nil
}

// Função que cadastra os herois com a devida normalização
func CadastrarHeroiComPoderesNormalizados(heroi Herois, idsPoderes []int) error {
	db := ConectaDB()
	defer db.Close()

	// Inicia uma transação para garantir consistência entre as tabelas
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	// Consulta para inserir o herói
	queryHeroi := `
		INSERT INTO Herois (
			nome_heroi, nome_real, sexo, altura, peso, data_nascimento, local_nascimento, 
			popularidade, forca, status_atividade
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id_heroi;
	`

	// Executa a consulta e captura o id do herói recém-inserido
	var idHeroi int
	err = tx.QueryRow(queryHeroi,
		heroi.NomeHeroi,
		heroi.Nome, // Nome real
		heroi.Sexo,
		heroi.Altura,
		heroi.Peso,
		heroi.DataNasc,
		heroi.LocalNasc,
		heroi.Popularidade,
		heroi.Forca,
		heroi.Status,
	).Scan(&idHeroi)

	if err != nil {
		tx.Rollback() // Reverte a transação em caso de erro
		return fmt.Errorf("erro ao cadastrar o herói: %w", err)
	}

	// Consulta para inserir na tabela herois_poderes
	queryHeroiPoder := `
		INSERT INTO herois_poderes (
			id_heroi, id_poder
		) VALUES ($1, $2);
	`

	// Itera sobre os IDs dos poderes e realiza as inserções na tabela herois_poderes
	for _, idPoder := range idsPoderes {
		_, err = tx.Exec(queryHeroiPoder, idHeroi, idPoder)
		if err != nil {
			tx.Rollback() // Reverte a transação em caso de erro
			return fmt.Errorf("erro ao associar herói e poder (id_heroi: %d, id_poder: %d): %w", idHeroi, idPoder, err)
		}
	}

	// Confirma a transação
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("erro ao confirmar transação: %w", err)
	}

	fmt.Println("Herói e associações com poderes cadastrados com sucesso!")
	return nil
}

// Função para remover um heroi(Oculta a exibição do heroi, porem mantem no BD)
func Remove(id int) error {
	log.Printf("Iniciando a remoção do herói com ID: %d", id)

	// Conexão com o banco de dados
	db := ConectaDB()
	defer func() {
		log.Println("Fechando a conexão com o banco de dados.")
		db.Close()
	}()

	// Query para deletar o herói
	query := `UPDATE herois SET hide = true WHERE id_heroi = $1`
	log.Printf("Executando a query: %s", query)

	// Executa a query
	result, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao executar a query: %v", err)
		return fmt.Errorf("erro ao remover herói com id %d: %w", id, err)
	}

	// Verifica se alguma linha foi afetada
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		return fmt.Errorf("erro ao verificar linhas afetadas ao remover herói com id %d: %w", id, err)
	}
	if rowsAffected == 0 {
		log.Printf("Nenhum herói encontrado com ID: %d", id)
		return fmt.Errorf("nenhum herói encontrado com id %d", id)
	}

	log.Printf("Herói com ID: %d removido com sucesso. Linhas afetadas: %d", id, rowsAffected)
	return nil
}

// Função para exibir todos os poderes
func ExibeTodosOsPoderes() []Poder {
	db := ConectaDB()
	defer db.Close()

	query := `SELECT * FROM Poderes ORDER BY id_poder ASC`
	TodosPoderes, err := db.Query(query)
	if err != nil {
		log.Fatalf("Erro ao executar a consulta: %v", err)
	}
	defer TodosPoderes.Close()

	var informacoes []Poder

	for TodosPoderes.Next() {
		var poder Poder
		err := TodosPoderes.Scan(
			&poder.Id_poder,
			&poder.Poder,
			&poder.Descricao,
		)
		if err != nil {
			log.Fatalf("Erro ao fazer o scan dos resultados: %v", err)
		}
		informacoes = append(informacoes, poder)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = TodosPoderes.Err(); err != nil {
		log.Fatalf("Erro durante a iteração dos resultados: %v", err)
	}
	return informacoes
}
