\connect siadmin_homo

DO
$$
BEGIN
    -- Inserir valores na tabela t_perfis apenas se eles ainda não existirem
    IF NOT EXISTS (SELECT 1 FROM t_perfis WHERE nome = 'Administrador de Sistema') THEN
        INSERT INTO t_perfis (nome, descricao, ativo, atualizado_em)
        VALUES ('Administrador de Sistema', 'Perfil de administrador com todas as permissões', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_perfis WHERE nome = 'Administrador de Operação') THEN
        INSERT INTO t_perfis (nome, descricao, ativo, atualizado_em)
        VALUES ('Administrador de Operação', 'Perfil de administrador com todas as permissões da parte de operação', TRUE, NOW());
    END IF;
    
    IF NOT EXISTS (SELECT 1 FROM t_perfis WHERE nome = 'Operação') THEN
        INSERT INTO t_perfis (nome, descricao, ativo, atualizado_em)
        VALUES ('Operação', 'Perfil de operação com apenas leitura das tabelas', TRUE, NOW());
    END IF;

    -- Inserir valores na tabela t_permissoes apenas se eles ainda não existirem

    -- Permissões relacionadas aos perfis

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-all-perfis') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-all-perfis', 'Permissão para buscar todos os perfis existentes no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-perfil-by-id') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-perfil-by-id', 'Permissão para buscar um perfil existente no banco de dados pelo Id', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuarios-vinculados') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuarios-vinculados', 'Permissão para buscar pelos usuários vinculados à um perfil pelo id do mesmo', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-permissoes-vinculadas') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-permissoes-vinculadas', 'Permissão para buscar pelas permissões vinculados à um perfil pelo id do mesmo', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'post-create-perfil') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('post-create-perfil', 'Permissão para criar um novo perfil no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'post-clone-perfil') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('post-clone-perfil', 'Permissão para clonar um perfil existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-update-perfil') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-update-perfil', 'Permissão para atualizar um perfil existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-restaurar-perfil') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-restaurar-perfil', 'Permissão para restaurar um perfil existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-desativar-perfil') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-desativar-perfil', 'Permissão para desativar um perfil existente no banco de dados', TRUE, NOW());
    END IF;

    -- Permissões relacionadas às permissões

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-all-permissoes') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-all-permissoes', 'Permissão para buscar por todas as permissões existentes no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-permissao-by-id') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-permissao-by-id', 'Permissão para buscar uma permissão existente no banco de dados pelo id', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'post-create-permissao') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('post-create-permissao', 'Permissão para criar uma nova permissão no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-update-permissao') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-update-permissao', 'Permissão para atualizar uma permissão existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-restaurar-permissao') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-restaurar-permissao', 'Permissão para restaurar uma permissão existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-desativar-permissao') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-desativar-permissao', 'Permissão para desativar uma permissão existente no banco de dados', TRUE, NOW());
    END IF;

    -- Permissões relacionadas aos usuários externos

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-all-usuarios-externos') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-all-usuarios-externos', 'Permissão para buscar por todos os usuários externos', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuario-externo-by-id') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuario-externo-by-id', 'Permissão para buscar um usuario-externo pelo Id', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuario-externo-by-id-externo') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuario-externo-by-id-externo', 'Permissão para buscar um usuario-externo pelo Id externo', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuario-externo-by-nome') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuario-externo-by-nome', 'Permissão para buscar um usuario-externo pelo nome', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuario-externo-by-documento') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuario-externo-by-documento', 'Permissão para buscar um usuario-externo pelo documento', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuario-externo-by-email') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuario-externo-by-email', 'Permissão para buscar um usuario-externo pelo e-mail', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuario-externo-by-codReserva') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuario-externo-by-codReserva', 'Permissão para buscar um usuario-externo pelo código de reserva', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'post-create-usuario-externo') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('post-create-usuario-externo', 'Permissão para criar um novo usuario-externo', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-update-usuario-externo') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-update-usuario-externo', 'Permissão para atualizar um usuario-externo existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-desativar-usuario-externo') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-desativar-usuario-externo', 'Permissão para desativar um usuario-externo', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-restaurar-usuario-externo') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-restaurar-usuario-externo', 'Permissão para restaurar usuarios-externos anteriormente desativados', TRUE, NOW());
    END IF;

    -- Permissões relacionadas aos usuários internos

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-all-usuarios-internos') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-all-usuarios-internos', 'Permissão para buscar todos os usuários existentes no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-usuario-interno-by-id') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-usuario-interno-by-id', 'Permissão para buscar um usuário existente no banco de dados pelo id', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'get-perfis-vinculados') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('get-perfis-vinculados', 'Permissão para buscar pelos perfis vinculados à um usuário pelo id do usuário', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'post-create-usuario-interno') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('post-create-usuario-interno', 'Permissão para criar um novo usuário no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'post-clone-usuario-interno') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('post-clone-usuario-interno', 'Permissão para clonar um usuário existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-update-usuario-interno') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-update-usuario-interno', 'Permissão para atualizar um usuário existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-alterar-senha-admin') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-alterar-senha-admin', 'Permissão de admin para alterar a senha de um usuário', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-alterar-propria-senha') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-alterar-propria-senha', 'Permissão de usuário para alterar a própria senha', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-restaurar-usuario-interno') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-restaurar-usuario-interno', 'Permissão para restaurar um usuário existente no banco de dados', TRUE, NOW());
    END IF;

    IF NOT EXISTS (SELECT 1 FROM t_permissoes WHERE nome = 'put-desativar-usuario-interno') THEN
        INSERT INTO t_permissoes (nome, descricao, ativo, atualizado_em)
        VALUES ('put-desativar-usuario-interno', 'Permissão para desativar um usuário existente no banco de dados', TRUE, NOW());
    END IF;
END
$$;
