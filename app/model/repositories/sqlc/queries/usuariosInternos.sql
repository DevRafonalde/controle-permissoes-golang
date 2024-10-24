-- name: FindAllUsuariosInternos :many
SELECT * FROM t_usuarios_internos WHERE id > $1 ORDER BY id LIMIT $2;

-- name: FindUsuarioInternoById :one
SELECT * FROM t_usuarios_internos WHERE id = $1;

-- name: FindUsuarioInternoByEmail :one
SELECT * FROM t_usuarios_internos WHERE email = $1;

-- name: CreateUsuarioInterno :one
INSERT INTO t_usuarios_internos (nome, email, senha, ativo, token_reset_senha, atualizado_em, senha_atualizada)
VALUES ($1, $2, $3, $4, $5, NOW(), $6)
RETURNING *;

-- name: UpdateUsuarioInterno :one
UPDATE t_usuarios_internos
SET nome = $2, email = $3, atualizado_em = NOW()
WHERE id = $1
RETURNING *;

-- name: DesativarUsuarioInterno :execrows
UPDATE t_usuarios_internos SET ativo = false WHERE id = $1
RETURNING id;

-- name: RestaurarUsuarioInterno :execrows
UPDATE t_usuarios_internos SET ativo = true WHERE id = $1
RETURNING id;

-- name: SetSenhaAtualizadaUsuarioInternoFalse :execrows
UPDATE t_usuarios_internos SET senha_atualizada = false WHERE id = $1
RETURNING id;

-- name: SetTokenResetSenhaUsuarioInterno :execrows
UPDATE t_usuarios_internos SET token_reset_senha = $1 WHERE id = $2
RETURNING id;

-- name: UpdateSenhaUsuarioInterno :execrows
UPDATE t_usuarios_internos SET senha = $1, senha_atualizada = true WHERE id = $2
RETURNING id;
