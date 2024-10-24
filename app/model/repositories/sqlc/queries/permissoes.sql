-- name: FindAllPermissoes :many
SELECT * FROM t_permissoes WHERE id > $1 ORDER BY id LIMIT $2;

-- name: FindPermissaoByID :one
SELECT * FROM t_permissoes WHERE id = $1;

-- name: FindPermissaoByNome :many
SELECT * FROM t_permissoes WHERE nome ILIKE '%' || $1 || '%' AND id > $2 ORDER BY id LIMIT $3;

-- name: CreatePermissao :one
INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
VALUES ($1, $2, true, NOW()) 
RETURNING *;

-- name: UpdatePermissao :one
UPDATE t_permissoes 
SET nome = $1, descricao = $2, ativo = $3, data_ultima_atualizacao = NOW()
WHERE id = $4
RETURNING *;
