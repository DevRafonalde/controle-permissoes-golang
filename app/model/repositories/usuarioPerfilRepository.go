package repositories

import (
	"context"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type UsuarioPerfilRepository interface {
	FindAll(context context.Context) ([]db.TUsuarioPerfil, error)
	FindByID(context context.Context, id int32) (db.TUsuarioPerfil, error)
	FindByPerfil(context context.Context, id int32) ([]db.TUsuarioPerfil, error)
	FindByUsuarioInterno(context context.Context, id int32) ([]db.TUsuarioPerfil, error)
	FindByUsuarioExterno(context context.Context, id int32) ([]db.TUsuarioPerfil, error)
	Create(context context.Context, usuarioPerfil db.CreateUsuarioPerfilParams) (db.TUsuarioPerfil, error)
	Update(context context.Context, usuarioPerfil db.UpdateUsuarioPerfilParams) (db.TUsuarioPerfil, error)
	Delete(context context.Context, id int32) error
}

type usuarioPerfilRepository struct {
	*db.Queries
}

func NewUsuarioPerfilRepository(queries *db.Queries) UsuarioPerfilRepository {
	return &usuarioPerfilRepository{Queries: queries}
}

func (usuarioPerfilRepository *usuarioPerfilRepository) FindAll(context context.Context) ([]db.TUsuarioPerfil, error) {
	usuarioPerfis, err := usuarioPerfilRepository.FindAllUsuarioPerfis(context)
	if err != nil {
		return nil, err
	}

	return usuarioPerfis, nil
}

func (usuarioPerfilRepository *usuarioPerfilRepository) FindByID(context context.Context, id int32) (db.TUsuarioPerfil, error) {
	usuarioPerfil, err := usuarioPerfilRepository.FindUsuarioPerfilByID(context, id)
	if err != nil {
		return db.TUsuarioPerfil{}, err
	}

	return usuarioPerfil, nil
}

func (usuarioPerfilRepository *usuarioPerfilRepository) FindByPerfil(context context.Context, perfilId int32) ([]db.TUsuarioPerfil, error) {
	usuarioPerfil, err := usuarioPerfilRepository.FindUsuarioPerfilByPerfil(context, perfilId)
	if err != nil {
		return []db.TUsuarioPerfil{}, err
	}

	return usuarioPerfil, nil
}

func (usuarioPerfilRepository *usuarioPerfilRepository) FindByUsuarioInterno(context context.Context, usuarioId int32) ([]db.TUsuarioPerfil, error) {
	usosUsuario, err := usuarioPerfilRepository.FindUsuarioPerfilByUsuarioInterno(context, pgtype.Int4{Int32: usuarioId})
	if err != nil {
		return []db.TUsuarioPerfil{}, err
	}

	return usosUsuario, nil
}

func (usuarioPerfilRepository *usuarioPerfilRepository) FindByUsuarioExterno(context context.Context, usuarioId int32) ([]db.TUsuarioPerfil, error) {
	usosUsuario, err := usuarioPerfilRepository.FindUsuarioPerfilByUsuarioExterno(context, pgtype.Int4{Int32: usuarioId})
	if err != nil {
		return []db.TUsuarioPerfil{}, err
	}

	return usosUsuario, nil
}

func (usuarioPerfilRepository *usuarioPerfilRepository) Create(context context.Context, usuarioPerfil db.CreateUsuarioPerfilParams) (db.TUsuarioPerfil, error) {
	usuarioPerfilCriado, err := usuarioPerfilRepository.CreateUsuarioPerfil(context, usuarioPerfil)
	if err != nil {
		return db.TUsuarioPerfil{}, err
	}

	return usuarioPerfilCriado, nil

}

func (usuarioPerfilRepository *usuarioPerfilRepository) Update(context context.Context, usuarioPerfil db.UpdateUsuarioPerfilParams) (db.TUsuarioPerfil, error) {
	usuarioPerfilAtualizado, err := usuarioPerfilRepository.UpdateUsuarioPerfil(context, usuarioPerfil)
	if err != nil {
		return db.TUsuarioPerfil{}, err
	}

	return usuarioPerfilAtualizado, nil
}

func (usuarioPerfilRepository *usuarioPerfilRepository) Delete(context context.Context, id int32) error {
	_, err := usuarioPerfilRepository.DeleteUsuarioPerfilById(context, id)
	return err
}
