package repositories

import (
	"context"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type PermissaoRepository interface {
	FindAll(context context.Context, cursor int32, tamanhoPagina int32) ([]db.TPermisso, error)
	FindByID(context context.Context, id int32) (db.TPermisso, error)
	FindByNome(context context.Context, nome string, cursor int32, tamanhoPagina int32) ([]db.TPermisso, error)
	Create(context context.Context, permissao db.CreatePermissaoParams) (db.TPermisso, error)
	Update(context context.Context, permissao db.UpdatePermissaoParams) (db.TPermisso, error)
}

type permissaoRepository struct {
	*db.Queries
}

func NewPermissaoRepository(queries *db.Queries) PermissaoRepository {
	return &permissaoRepository{Queries: queries}
}

func (permissaoRepository *permissaoRepository) FindAll(context context.Context, cursor int32, tamanhoPagina int32) ([]db.TPermisso, error) {
	permissoes, err := permissaoRepository.FindAllPermissoes(context, db.FindAllPermissoesParams{ID: cursor, Limit: tamanhoPagina})
	if err != nil {
		return nil, err
	}

	return permissoes, nil
}

func (permissaoRepository *permissaoRepository) FindByID(context context.Context, id int32) (db.TPermisso, error) {
	permissao, err := permissaoRepository.FindPermissaoByID(context, id)
	if err != nil {
		return db.TPermisso{}, err
	}

	return permissao, nil
}

func (permissaoRepository *permissaoRepository) FindByNome(context context.Context, nome string, cursor int32, tamanhoPagina int32) ([]db.TPermisso, error) {
	permissao, err := permissaoRepository.FindPermissaoByNome(context, db.FindPermissaoByNomeParams{Column1: pgtype.Text{String: nome, Valid: true}, ID: cursor, Limit: tamanhoPagina})
	if err != nil {
		return []db.TPermisso{}, err
	}

	return permissao, nil
}

func (permissaoRepository *permissaoRepository) Create(context context.Context, permissao db.CreatePermissaoParams) (db.TPermisso, error) {
	permissaoCriada, err := permissaoRepository.CreatePermissao(context, permissao)
	if err != nil {
		return db.TPermisso{}, err
	}

	return permissaoCriada, nil
}

func (permissaoRepository *permissaoRepository) Update(context context.Context, permissao db.UpdatePermissaoParams) (db.TPermisso, error) {
	permissaoAtualizada, err := permissaoRepository.UpdatePermissao(context, permissao)
	if err != nil {
		return db.TPermisso{}, err
	}

	return permissaoAtualizada, nil
}
