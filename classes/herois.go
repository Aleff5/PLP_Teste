package classes

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"teste/database"
	"time"
)

// Estrutura das informações pessoais
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

// Método para exibir todos os heróis
func ExibeTodosOsNomes() []string {
	db := database.ConectaDB()
	defer db.Close()

	query := `SELECT nome_heroi FROM Herois WHERE esconder = false`
	TodosHerois, err := db.Query(query)
	if err != nil {
		log.Fatalf("Erro ao executar a consulta: %v", err)
	}
	defer TodosHerois.Close()

	var informacoes []string

	for TodosHerois.Next() {
		var heroi string
		err := TodosHerois.Scan(
			&heroi,
		)
		if err != nil {
			log.Fatalf("Erro ao fazer o scan dos resultados: %v", err)
		}
		informacoes = append(informacoes, heroi)
	}
	// Verifica se ocorreu algum erro durante a iteração
	if err = TodosHerois.Err(); err != nil {
		log.Fatalf("Erro durante a iteração dos resultados: %v", err)
	}
	return informacoes
}

// Método para exibir as informações gerais dos heróis
func (h Herois) ExibeInfosGerais() []Herois {

	db := database.ConectaDB()
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
			h.esconder = false
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
	db := database.ConectaDB()
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
			h.esconder = false
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
// Método para exibir os nomes dos heróis por popularidade
func BuscaHeroisPorPopularidade(popularidade int) ([]string, error) {
	db := database.ConectaDB()
	defer db.Close()

	query := `
        SELECT 
            h.nome_heroi
        FROM 
            Herois h
        WHERE 
            h.popularidade <= $1
        AND
            h.esconder = false
    `

	rows, err := db.Query(query, popularidade)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nomesHerois []string
	for rows.Next() {
		var nomeHeroi string
		if err := rows.Scan(&nomeHeroi); err != nil {
			return nil, err
		}
		nomesHerois = append(nomesHerois, nomeHeroi)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nomesHerois, nil
}

// Método para exibir as informações dos heróis por status
// Método para exibir os nomes dos heróis por status
func BuscaHeroisPorStatus(status string) ([]string, error) {
	db := database.ConectaDB()
	defer db.Close()

	query := `
        SELECT 
            h.nome_heroi
        FROM 
            Herois h
        WHERE 
            h.status_atividade = $1
        AND
            h.esconder = false
    `

	rows, err := db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nomesHerois []string
	for rows.Next() {
		var nomeHeroi string
		if err := rows.Scan(&nomeHeroi); err != nil {
			return nil, err
		}
		nomesHerois = append(nomesHerois, nomeHeroi)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nomesHerois, nil
}

// Função que cadastra os herois com a devida normalização
func CadastrarHeroiComPoderesNormalizados(heroi Herois, idsPoderes []int) error {
	db := database.ConectaDB()
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
func Remove(nome string) error {
	log.Printf("Iniciando a remoção do herói: %s", nome)

	// Conexão com o banco de dados
	db := database.ConectaDB()
	defer func() {
		log.Println("Fechando a conexão com o banco de dados.")
		db.Close()
	}()

	// Query para deletar o herói
	query := `UPDATE herois SET esconder = true WHERE nome_heroi = $1`
	log.Printf("Executando a query: %s", query)

	// Executa a query
	result, err := db.Exec(query, nome)
	if err != nil {
		log.Printf("Erro ao executar a query: %v", err)
		return fmt.Errorf("erro ao remover herói com id %s: %w", nome, err)
	}

	// Verifica se alguma linha foi afetada
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		return fmt.Errorf("erro ao verificar linhas afetadas ao remover herói com id %s: %w", nome, err)
	}
	if rowsAffected == 0 {
		log.Printf("Nenhum herói encontrado com ID: %s", nome)
		return fmt.Errorf("nenhum herói encontrado com id %s", nome)
	}

	log.Printf("Herói: %s removido com sucesso. Linhas afetadas: %d", nome, rowsAffected)
	return nil
}

// Função para exibir todos os poderes
func ExibeTodosOsPoderes() []Poder {
	db := database.ConectaDB()
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

// Função para Modificar Informações do Herói
func EditarHeroiPorNome(nomeHeroi string, heroiAtualizado Herois) error {
	db := database.ConectaDB()
	defer db.Close()

	// Consulta para atualizar os dados do herói com base no nome
	query := `
        UPDATE Herois
        SET 
            nome_heroi = $1,
            nome_real = $2,
            sexo = $3,
            altura = $4,
            peso = $5,
            data_nascimento = $6,
            local_nascimento = $7,
            popularidade = $8,
            forca = $9,
            status_atividade = $10
        WHERE nome_heroi = $11;
	`

	// Executa a consulta com os valores atualizados
	_, err := db.Exec(query,
		heroiAtualizado.NomeHeroi, // Atualiza o nome do herói
		heroiAtualizado.Nome,      // Nome real
		heroiAtualizado.Sexo,
		heroiAtualizado.Altura,
		heroiAtualizado.Peso,
		heroiAtualizado.DataNasc,
		heroiAtualizado.LocalNasc,
		heroiAtualizado.Popularidade,
		heroiAtualizado.Forca,
		heroiAtualizado.Status, // Status da atividade
		nomeHeroi,              // Nome atual do herói para identificar o registro
	)

	if err != nil {
		return fmt.Errorf("erro ao editar herói com nome '%s': %w", nomeHeroi, err)
	}

	fmt.Printf("Herói com nome '%s' atualizado com sucesso!\n", nomeHeroi)
	return nil
}