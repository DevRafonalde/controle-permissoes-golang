-- name: FindAllPerfis :many
SELECT * FROM t_perfis WHERE id > $1 ORDER BY id LIMIT $2;

-- name: FindPerfilByID :one
SELECT * FROM t_perfis WHERE id = $1;

-- name: FindPerfilByNome :many
SELECT * FROM t_perfis WHERE nome ILIKE '%' || $1 || '%' AND id > $2 ORDER BY id LIMIT $3;

-- name: FindPerfilByPermissao :many
SELECT p.* 
FROM t_perfis p
JOIN t_perfil_permissao pp ON p.id = pp.perfil_id
WHERE pp.permissao_id = $1;

-- name: CreatePerfil :one
INSERT INTO t_perfis (nome, descricao, ativo, atualizado_em)
VALUES ($1, $2, true, NOW())
RETURNING *;

-- name: UpdatePerfil :one
UPDATE t_perfis 
SET nome = $1, descricao = $2, ativo = $3, atualizado_em = NOW()
WHERE id = $4
RETURNING *;
