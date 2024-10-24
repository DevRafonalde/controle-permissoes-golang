-- name: FindAllUsuarioPerfis :many
SELECT * FROM t_usuario_perfil;

-- name: FindUsuarioPerfilByID :one
SELECT * FROM t_usuario_perfil WHERE id = $1;

-- name: FindUsuarioPerfilByPerfil :many
SELECT * FROM t_usuario_perfil WHERE perfil_id = $1;

-- name: FindUsuarioPerfilByUsuarioInterno :many
SELECT * FROM t_usuario_perfil WHERE usuario_interno_id = $1;

-- name: FindUsuarioPerfilByUsuarioExterno :many
SELECT * FROM t_usuario_perfil WHERE usuario_externo_id = $1;

-- name: CreateUsuarioPerfil :one
INSERT INTO t_usuario_perfil (usuario_interno_id, usuario_externo_id, perfil_id, data_hora)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUsuarioPerfil :one
UPDATE t_usuario_perfil 
SET usuario_interno_id = $1, usuario_externo_id = $2, perfil_id = $3, data_hora = $4
WHERE id = $5
RETURNING *;

-- name: DeleteUsuarioPerfilById :execrows
DELETE FROM t_usuario_perfil 
WHERE id = $1 
RETURNING *;
