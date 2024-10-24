-- name: FindAllPerfilPermissao :many
SELECT * FROM t_perfil_permissao;

-- name: FindPerfilPermissaoByID :one
SELECT * FROM t_perfil_permissao WHERE id = $1;

-- name: FindPerfilPermissaoByPerfil :many
SELECT * FROM t_perfil_permissao WHERE perfil_id = $1;

-- name: FindPerfilPermissaoByPermissao :many
SELECT * FROM t_perfil_permissao WHERE permissao_id = $1;

-- name: CreatePerfilPermissao :one
INSERT INTO t_perfil_permissao (perfil_id, permissao_id, data_hora) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdatePerfilPermissao :one
UPDATE t_perfil_permissao 
SET perfil_id = $1, permissao_id = $2, data_hora = $3
WHERE id = $4
RETURNING *;

-- name: DeletePerfilPermissaoById :execrows
DELETE FROM t_perfil_permissao WHERE id = $1 
RETURNING id;
