package repositories

import (
	"context"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type UsuarioExternoRepository interface {
	FindAll(context context.Context, cursor int32, tamanhoPagina int32) ([]db.TUsuariosExterno, error)
	FindByID(context context.Context, id int32) (db.TUsuariosExterno, error)
	FindByIDExterno(context context.Context, id int32) (db.TUsuariosExterno, error)
	FindByNome(context context.Context, nome string, cursor int32, tamanhoPagina int32) ([]db.TUsuariosExterno, error)
	FindByDocumento(context context.Context, documento string) (db.TUsuariosExterno, error)
	FindByEmail(context context.Context, email string) (db.TUsuariosExterno, error)
	FindByCodReserva(context context.Context, codIbge string) (db.TUsuariosExterno, error)
	Create(context context.Context, cliente db.CreateUsuarioExternoParams) (db.TUsuariosExterno, error)
	Update(context context.Context, cliente db.UpdateUsuarioExternoParams) (db.TUsuariosExterno, error)
	SetSenhaAtualizada(context context.Context, id int32) (int64, error)
	Desativar(context context.Context, id int32) (int64, error)
	Restaurar(context context.Context, id int32) (int64, error)
	UpdateSenha(context context.Context, senha string, id int32) (int64, error)
	SetTokenResetSenha(context context.Context, token string, id int32) (int64, error)
}

type usuarioExternoRepository struct {
	*db.Queries
}

func NewUsuarioExternoRepository(queries *db.Queries) UsuarioExternoRepository {
	return &usuarioExternoRepository{
		Queries: queries,
	}
}

func (usuarioExternoRepository *usuarioExternoRepository) FindAll(ctx context.Context, cursor int32, tamanhoPagina int32) ([]db.TUsuariosExterno, error) {
	return usuarioExternoRepository.FindAllUsuariosExternos(ctx, db.FindAllUsuariosExternosParams{ID: cursor, Limit: tamanhoPagina})

}

func (usuarioExternoRepository *usuarioExternoRepository) FindByID(context context.Context, id int32) (db.TUsuariosExterno, error) {
	cliente, err := usuarioExternoRepository.FindUsuarioExternoByID(context, id)
	if err != nil {
		return db.TUsuariosExterno{}, err
	}

	return cliente, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) FindByIDExterno(context context.Context, id int32) (db.TUsuariosExterno, error) {
	cliente, err := usuarioExternoRepository.FindUsuarioExternoByIDExterno(context, id)
	if err != nil {
		return db.TUsuariosExterno{}, err
	}

	return cliente, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) FindByNome(context context.Context, nome string, cursor int32, tamanhoPagina int32) ([]db.TUsuariosExterno, error) {
	cliente, err := usuarioExternoRepository.FindUsuarioExternoByNome(context, db.FindUsuarioExternoByNomeParams{Column1: pgtype.Text{String: nome, Valid: true}, ID: cursor, Limit: tamanhoPagina})
	if err != nil {
		return []db.TUsuariosExterno{}, err
	}

	return cliente, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) FindByDocumento(context context.Context, documento string) (db.TUsuariosExterno, error) {
	clientes, err := usuarioExternoRepository.FindUsuarioExternoByDocumento(context, documento)
	if err != nil {
		return db.TUsuariosExterno{}, err
	}

	return clientes, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) FindByEmail(context context.Context, email string) (db.TUsuariosExterno, error) {
	clientes, err := usuarioExternoRepository.FindUsuarioExternoByEmail(context, email)
	if err != nil {
		return db.TUsuariosExterno{}, err
	}

	return clientes, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) FindByCodReserva(context context.Context, codIbge string) (db.TUsuariosExterno, error) {
	cliente, err := usuarioExternoRepository.FindUsuarioExternoByCodReserva(context, pgtype.Text{String: codIbge, Valid: true})
	if err != nil {
		return db.TUsuariosExterno{}, err
	}

	return cliente, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) Create(context context.Context, cliente db.CreateUsuarioExternoParams) (db.TUsuariosExterno, error) {
	clienteCriada, err := usuarioExternoRepository.CreateUsuarioExterno(context, cliente)
	if err != nil {
		return db.TUsuariosExterno{}, err
	}

	return clienteCriada, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) Update(context context.Context, cliente db.UpdateUsuarioExternoParams) (db.TUsuariosExterno, error) {
	clienteAtualizada, err := usuarioExternoRepository.UpdateUsuarioExterno(context, cliente)
	if err != nil {
		return db.TUsuariosExterno{}, err
	}

	return clienteAtualizada, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) SetSenhaAtualizada(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := usuarioExternoRepository.SetSenhaAtualizadaUsuarioExternoFalse(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) Desativar(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := usuarioExternoRepository.DesativarUsuarioExterno(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) Restaurar(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := usuarioExternoRepository.RestaurarUsuarioExterno(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) UpdateSenha(context context.Context, senha string, id int32) (int64, error) {
	linhasAfetadas, err := usuarioExternoRepository.UpdateSenhaUsuarioExterno(context, db.UpdateSenhaUsuarioExternoParams{Senha: senha, ID: id})
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}

func (usuarioExternoRepository *usuarioExternoRepository) SetTokenResetSenha(context context.Context, token string, id int32) (int64, error) {
	linhasAfetadas, err := usuarioExternoRepository.SetTokenResetSenhaUsuarioExterno(context, db.SetTokenResetSenhaUsuarioExternoParams{TokenResetSenha: pgtype.Text{String: token, Valid: true}, ID: id})
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
