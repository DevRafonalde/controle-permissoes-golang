-- Cria tabela de usuários externos se ela não existir
CREATE TABLE IF NOT EXISTS public.t_usuarios_externos (
	id SERIAL PRIMARY KEY,
	"uuid" varchar(255) NULL UNIQUE,
	id_externo int NOT NULL UNIQUE,
	nome varchar(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
	senha VARCHAR(255) NOT NULL,
	documento varchar(14) NOT NULL UNIQUE,
	cod_reserva varchar(12) NULL,
	ativo BOOLEAN,
	atualizado_em timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
    token_reset_senha VARCHAR(255),
	senha_atualizada BOOLEAN
);

-- Cria tabela de usuários se ela não existir
CREATE TABLE IF NOT EXISTS t_usuarios_internos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
	senha VARCHAR(255) NOT NULL,
    ativo BOOLEAN,
	token_reset_senha VARCHAR(255),
    atualizado_em timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
	senha_atualizada BOOLEAN
);

-- Cria tabela de perfis se ela não existir
CREATE TABLE IF NOT EXISTS t_perfis (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    ativo BOOLEAN,
    atualizado_em timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);

-- Cria tabela de permissões se ela não existir
CREATE TABLE IF NOT EXISTS t_permissoes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) UNIQUE NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    ativo BOOLEAN,
    atualizado_em timestamptz DEFAULT CURRENT_TIMESTAMP NULL
);

-- Cria tabela de associação de usuários a perfis se ela não existir
CREATE TABLE IF NOT EXISTS t_usuario_perfil (
    id SERIAL PRIMARY KEY,
    usuario_interno_id INT NULL,
    usuario_externo_id INT NULL,
    perfil_id INT NOT NULL,
    data_hora TIMESTAMP,
    FOREIGN KEY (usuario_interno_id) REFERENCES t_usuarios_internos(id),
    FOREIGN KEY (usuario_externo_id) REFERENCES t_usuarios_externos(id),
    FOREIGN KEY (perfil_id) REFERENCES t_perfis(id),
	UNIQUE (usuario_interno_id, perfil_id), -- Garante que um usuário não pode estar ligado ao mesmo perfil mais de uma vez
    UNIQUE (usuario_externo_id, perfil_id)
);

-- Cria tabela de associação de perfis a permissões se ela não existir
CREATE TABLE IF NOT EXISTS t_perfil_permissao (
    id SERIAL PRIMARY KEY,
    perfil_id INT NOT NULL,
    permissao_id INT NOT NULL,
    data_hora TIMESTAMP,
    FOREIGN KEY (perfil_id) REFERENCES t_perfis(id),
    FOREIGN KEY (permissao_id) REFERENCES t_permissoes(id),
	UNIQUE (perfil_id, permissao_id) -- Garante que o um perfil não pode estar ligado à mesma permissão mais de uma vez
);
