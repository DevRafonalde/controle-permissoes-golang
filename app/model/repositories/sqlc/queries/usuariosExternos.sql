-- name: FindAllUsuariosExternos :many
SELECT * FROM t_usuarios_externos WHERE id > $1 ORDER BY id LIMIT $2;

-- name: FindUsuarioExternoByID :one
SELECT * FROM t_usuarios_externos WHERE id = $1;

-- name: FindUsuarioExternoByIDExterno :one
SELECT * FROM t_usuarios_externos WHERE id_externo = $1;

-- name: FindUsuarioExternoByNome :many
SELECT * FROM t_usuarios_externos WHERE nome ILIKE '%' || $1 || '%' AND id > $2 ORDER BY id LIMIT $3;

-- name: FindUsuarioExternoByDocumento :one
SELECT * FROM t_usuarios_externos WHERE documento = $1;

-- name: FindUsuarioExternoByEmail :one
SELECT * FROM t_usuarios_externos WHERE email = $1;

-- name: FindUsuarioExternoByCodReserva :one
SELECT * FROM t_usuarios_externos WHERE cod_reserva = $1;

-- name: CreateUsuarioExterno :one
INSERT INTO t_usuarios_externos (uuid, id_externo, nome, email, senha, documento, cod_reserva, ativo, atualizado_em)
VALUES ($1, $2, $3, $4, $5, $6, $7, true, NOW())
RETURNING *;

-- name: UpdateUsuarioExterno :one
UPDATE t_usuarios_externos 
SET uuid = $1, id_externo = $2, nome = $3, email = $4, senha = $5, documento = $6, cod_reserva = $7, atualizado_em = NOW()
WHERE id = $8
RETURNING *;

-- name: DesativarUsuarioExterno :execrows
UPDATE t_usuarios_externos SET ativo = false WHERE id = $1 
RETURNING id;

-- name: RestaurarUsuarioExterno :execrows
UPDATE t_usuarios_externos SET ativo = true WHERE id = $1 
RETURNING id;

-- name: SetSenhaAtualizadaUsuarioExternoFalse :execrows
UPDATE t_usuarios_externos SET senha_atualizada = false WHERE id = $1
RETURNING id;

-- name: SetTokenResetSenhaUsuarioExterno :execrows
UPDATE t_usuarios_externos SET token_reset_senha = $1 WHERE id = $2
RETURNING id;

-- name: UpdateSenhaUsuarioExterno :execrows
UPDATE t_usuarios_externos SET senha = $1, senha_atualizada = true WHERE id = $2
RETURNING id;
