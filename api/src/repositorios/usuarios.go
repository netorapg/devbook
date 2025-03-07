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
func (u usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
