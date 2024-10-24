package repositories

import (
	"context"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type UsuarioInternoRepository interface {
	FindAll(context context.Context, cursor int32, tamanhoPagina int32) ([]db.TUsuariosInterno, error)
	FindByID(context context.Context, id int32) (db.TUsuariosInterno, error)
	FindByEmail(context context.Context, email string) (db.TUsuariosInterno, error)
	Create(context context.Context, usuario db.CreateUsuarioInternoParams) (db.TUsuariosInterno, error)
	Update(context context.Context, usuario db.UpdateUsuarioInternoParams) (db.TUsuariosInterno, error)
	SetSenhaAtualizada(context context.Context, id int32) (int64, error)
	Desativar(context context.Context, id int32) (int64, error)
	Restaurar(context context.Context, id int32) (int64, error)
	UpdateSenha(context context.Context, senha string, id int32) (int64, error)
	SetTokenResetSenha(context context.Context, token string, id int32) (int64, error)
}

type usuarioInternoRepository struct {
	*db.Queries
}

func NewUsuarioInternoRepository(queries *db.Queries) UsuarioInternoRepository {
	return &usuarioInternoRepository{Queries: queries}
}

func (usuarioInternoRepository *usuarioInternoRepository) FindAll(context context.Context, cursor int32, tamanhoPagina int32) ([]db.TUsuariosInterno, error) {
	usuarios, err := usuarioInternoRepository.FindAllUsuariosInternos(context, db.FindAllUsuariosInternosParams{ID: cursor, Limit: tamanhoPagina})
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) FindByID(context context.Context, id int32) (db.TUsuariosInterno, error) {
	usuario, err := usuarioInternoRepository.FindUsuarioInternoById(context, id)
	if err != nil {
		return db.TUsuariosInterno{}, err
	}

	return usuario, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) FindByEmail(context context.Context, email string) (db.TUsuariosInterno, error) {
	usuario, err := usuarioInternoRepository.FindUsuarioInternoByEmail(context, email)
	if err != nil {
		return db.TUsuariosInterno{}, err
	}

	return usuario, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) Create(context context.Context, usuario db.CreateUsuarioInternoParams) (db.TUsuariosInterno, error) {
	usuarioCriado, err := usuarioInternoRepository.CreateUsuarioInterno(context, usuario)
	if err != nil {
		return db.TUsuariosInterno{}, err
	}

	return usuarioCriado, nil

}

func (usuarioInternoRepository *usuarioInternoRepository) Update(context context.Context, usuario db.UpdateUsuarioInternoParams) (db.TUsuariosInterno, error) {
	usuarioAtualizado, err := usuarioInternoRepository.UpdateUsuarioInterno(context, usuario)
	if err != nil {
		return db.TUsuariosInterno{}, err
	}

	return usuarioAtualizado, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) SetSenhaAtualizada(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := usuarioInternoRepository.SetSenhaAtualizadaUsuarioInternoFalse(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) Desativar(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := usuarioInternoRepository.DesativarUsuarioInterno(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) Restaurar(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := usuarioInternoRepository.RestaurarUsuarioInterno(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) UpdateSenha(context context.Context, senha string, id int32) (int64, error) {
	linhasAfetadas, err := usuarioInternoRepository.UpdateSenhaUsuarioInterno(context, db.UpdateSenhaUsuarioInternoParams{Senha: senha, ID: id})
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioInternoRepository *usuarioInternoRepository) SetTokenResetSenha(context context.Context, token string, id int32) (int64, error) {
	linhasAfetadas, err := usuarioInternoRepository.SetTokenResetSenhaUsuarioInterno(context, db.SetTokenResetSenhaUsuarioInternoParams{TokenResetSenha: pgtype.Text{String: token, Valid: true}, ID: id})
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
