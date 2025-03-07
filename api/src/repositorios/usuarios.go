package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Cria e insere um novo usuário no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	// Prepara a instrução SQL com RETURNING para obter o ID gerado
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values ($1, $2, $3, $4) RETURNING id",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	// Executa a inserção e captura o ID retornado
	var ultimoIDInserido uint64
	err := statement.QueryRow(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha).Scan(&ultimoIDInserido)
	if err != nil {
		return 0, err
	}

	return ultimoIDInserido, nil
}
