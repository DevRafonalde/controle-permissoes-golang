// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: perfilPermissao.sql

package repositoryIMPL

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPerfilPermissao = `-- name: CreatePerfilPermissao :one
INSERT INTO t_perfil_permissao (perfil_id, permissao_id, data_hora) 
VALUES ($1, $2, $3)
RETURNING id, perfil_id, permissao_id, data_hora
`

type CreatePerfilPermissaoParams struct {
	PerfilID    int32
	PermissaoID int32
	DataHora    pgtype.Timestamp
}

func (q *Queries) CreatePerfilPermissao(ctx context.Context, arg CreatePerfilPermissaoParams) (TPerfilPermissao, error) {
	row := q.db.QueryRow(ctx, createPerfilPermissao, arg.PerfilID, arg.PermissaoID, arg.DataHora)
	var i TPerfilPermissao
	err := row.Scan(
		&i.ID,
		&i.PerfilID,
		&i.PermissaoID,
		&i.DataHora,
	)
	return i, err
}

const deletePerfilPermissaoById = `-- name: DeletePerfilPermissaoById :execrows
DELETE FROM t_perfil_permissao WHERE id = $1 
RETURNING id
`

func (q *Queries) DeletePerfilPermissaoById(ctx context.Context, id int32) (int64, error) {
	result, err := q.db.Exec(ctx, deletePerfilPermissaoById, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAllPerfilPermissao = `-- name: FindAllPerfilPermissao :many
SELECT id, perfil_id, permissao_id, data_hora FROM t_perfil_permissao
`

func (q *Queries) FindAllPerfilPermissao(ctx context.Context) ([]TPerfilPermissao, error) {
	rows, err := q.db.Query(ctx, findAllPerfilPermissao)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TPerfilPermissao
	for rows.Next() {
		var i TPerfilPermissao
		if err := rows.Scan(
			&i.ID,
			&i.PerfilID,
			&i.PermissaoID,
			&i.DataHora,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findPerfilPermissaoByID = `-- name: FindPerfilPermissaoByID :one
SELECT id, perfil_id, permissao_id, data_hora FROM t_perfil_permissao WHERE id = $1
`

func (q *Queries) FindPerfilPermissaoByID(ctx context.Context, id int32) (TPerfilPermissao, error) {
	row := q.db.QueryRow(ctx, findPerfilPermissaoByID, id)
	var i TPerfilPermissao
	err := row.Scan(
		&i.ID,
		&i.PerfilID,
		&i.PermissaoID,
		&i.DataHora,
	)
	return i, err
}

const findPerfilPermissaoByPerfil = `-- name: FindPerfilPermissaoByPerfil :many
SELECT id, perfil_id, permissao_id, data_hora FROM t_perfil_permissao WHERE perfil_id = $1
`

func (q *Queries) FindPerfilPermissaoByPerfil(ctx context.Context, perfilID int32) ([]TPerfilPermissao, error) {
	rows, err := q.db.Query(ctx, findPerfilPermissaoByPerfil, perfilID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TPerfilPermissao
	for rows.Next() {
		var i TPerfilPermissao
		if err := rows.Scan(
			&i.ID,
			&i.PerfilID,
			&i.PermissaoID,
			&i.DataHora,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findPerfilPermissaoByPermissao = `-- name: FindPerfilPermissaoByPermissao :many
SELECT id, perfil_id, permissao_id, data_hora FROM t_perfil_permissao WHERE permissao_id = $1
`

func (q *Queries) FindPerfilPermissaoByPermissao(ctx context.Context, permissaoID int32) ([]TPerfilPermissao, error) {
	rows, err := q.db.Query(ctx, findPerfilPermissaoByPermissao, permissaoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TPerfilPermissao
	for rows.Next() {
		var i TPerfilPermissao
		if err := rows.Scan(
			&i.ID,
			&i.PerfilID,
			&i.PermissaoID,
			&i.DataHora,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePerfilPermissao = `-- name: UpdatePerfilPermissao :one
UPDATE t_perfil_permissao 
SET perfil_id = $1, permissao_id = $2, data_hora = $3
WHERE id = $4
RETURNING id, perfil_id, permissao_id, data_hora
`

type UpdatePerfilPermissaoParams struct {
	PerfilID    int32
	PermissaoID int32
	DataHora    pgtype.Timestamp
	ID          int32
}

func (q *Queries) UpdatePerfilPermissao(ctx context.Context, arg UpdatePerfilPermissaoParams) (TPerfilPermissao, error) {
	row := q.db.QueryRow(ctx, updatePerfilPermissao,
		arg.PerfilID,
		arg.PermissaoID,
		arg.DataHora,
		arg.ID,
	)
	var i TPerfilPermissao
	err := row.Scan(
		&i.ID,
		&i.PerfilID,
		&i.PermissaoID,
		&i.DataHora,
	)
	return i, err
}