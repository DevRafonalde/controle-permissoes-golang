package repositories

import (
	"context"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"
)

type PerfilPermissaoRepository interface {
	FindAll(context context.Context) ([]db.TPerfilPermissao, error)
	FindByID(context context.Context, id int32) (db.TPerfilPermissao, error)
	FindByPerfil(context context.Context, perfilId int32) ([]db.TPerfilPermissao, error)
	FindByPermissao(context context.Context, permissaoId int32) ([]db.TPerfilPermissao, error)
	Create(context context.Context, perfilPermissao db.CreatePerfilPermissaoParams) (db.TPerfilPermissao, error)
	Update(context context.Context, perfilPermissao db.UpdatePerfilPermissaoParams) (db.TPerfilPermissao, error)
	Delete(context context.Context, id int32) error
}

type perfilPermissaoRepository struct {
	*db.Queries
}

func NewPerfilPermissaoRepository(queries *db.Queries) PerfilPermissaoRepository {
	return &perfilPermissaoRepository{Queries: queries}
}

func (perfilPermissaoRepository *perfilPermissaoRepository) FindAll(context context.Context) ([]db.TPerfilPermissao, error) {
	perfisPermissao, err := perfilPermissaoRepository.FindAllPerfilPermissao(context)
	if err != nil {
		return nil, err
	}

	return perfisPermissao, nil
}

func (perfilPermissaoRepository *perfilPermissaoRepository) FindByID(context context.Context, id int32) (db.TPerfilPermissao, error) {
	perfilPermissao, err := perfilPermissaoRepository.FindPerfilPermissaoByID(context, id)
	if err != nil {
		return db.TPerfilPermissao{}, err
	}

	return perfilPermissao, nil
}

func (perfilPermissaoRepository *perfilPermissaoRepository) FindByPerfil(context context.Context, perfilId int32) ([]db.TPerfilPermissao, error) {
	perfisPermissao, err := perfilPermissaoRepository.FindPerfilPermissaoByPerfil(context, perfilId)
	if err != nil {
		return []db.TPerfilPermissao{}, err
	}

	return perfisPermissao, nil
}

func (perfilPermissaoRepository *perfilPermissaoRepository) FindByPermissao(context context.Context, permissaoId int32) ([]db.TPerfilPermissao, error) {
	usosPermissao, err := perfilPermissaoRepository.FindPerfilPermissaoByPermissao(context, permissaoId)
	if err != nil {
		return []db.TPerfilPermissao{}, err
	}

	return usosPermissao, nil
}

func (perfilPermissaoRepository *perfilPermissaoRepository) Create(context context.Context, perfilPermissao db.CreatePerfilPermissaoParams) (db.TPerfilPermissao, error) {
	perfilPermissaoCriado, err := perfilPermissaoRepository.CreatePerfilPermissao(context, perfilPermissao)
	if err != nil {
		return db.TPerfilPermissao{}, err
	}

	return perfilPermissaoCriado, nil

}

func (perfilPermissaoRepository *perfilPermissaoRepository) Update(context context.Context, perfilPermissao db.UpdatePerfilPermissaoParams) (db.TPerfilPermissao, error) {
	perfilPermissaoAtualizado, err := perfilPermissaoRepository.UpdatePerfilPermissao(context, perfilPermissao)
	if err != nil {
		return db.TPerfilPermissao{}, err
	}

	return perfilPermissaoAtualizado, nil
}

func (perfilPermissaoRepository *perfilPermissaoRepository) Delete(context context.Context, id int32) error {
	_, err := perfilPermissaoRepository.DeletePerfilPermissaoById(context, id)
	if err != nil {
		return err
	}

	return nil
}
