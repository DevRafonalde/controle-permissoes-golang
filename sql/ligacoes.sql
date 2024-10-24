\connect siadmin_homo

DO $$
DECLARE
    perm RECORD;
BEGIN
    -- Ligando todas as permissões ao perfil de id 1
    FOR perm IN (SELECT id FROM t_permissoes) LOOP
        IF NOT EXISTS (
            SELECT 1
            FROM t_perfil_permissao
            WHERE perfil_id = 1 AND permissao_id = perm.id
            ) THEN
            INSERT INTO t_perfil_permissao (perfil_id, permissao_id, data_hora)
            VALUES (1, perm.id, NOW());
        END IF;
    END LOOP;

    -- Ligando todas as permissões que não estão relacionadas a usuários, perfis ou permissões ao perfil de id 2
    FOR perm IN (
        SELECT id
        FROM t_permissoes
        WHERE nome NOT LIKE '%usuario%'
        AND nome NOT LIKE '%usuarios%'
        AND nome NOT LIKE '%perfil%'
        AND nome NOT LIKE '%perfis%'
        AND nome NOT LIKE '%permissao%'
        AND nome NOT LIKE '%permissoes%'
    ) LOOP
        IF NOT EXISTS (
            SELECT 1 FROM t_perfil_permissao
            WHERE perfil_id = 2
            AND permissao_id = perm.id
        ) THEN
            INSERT INTO t_perfil_permissao (perfil_id, permissao_id, data_hora)
            VALUES (2, perm.id, NOW());
        END IF;
    END LOOP;

    -- Ligando todas as permissões que possuem "get" ao perfil de id 3
    FOR perm IN (
        SELECT id FROM t_permissoes
        WHERE nome LIKE 'get-%'
        AND nome NOT LIKE '%usuario%'
        AND nome NOT LIKE '%usuarios%'
        AND nome NOT LIKE '%perfil%'
        AND nome NOT LIKE '%perfis%'
        AND nome NOT LIKE '%permissao%'
        AND nome NOT LIKE '%permissoes%'
        OR nome LIKE '%propria%'
    ) LOOP
        IF NOT EXISTS (
            SELECT 1
            FROM t_perfil_permissao
            WHERE perfil_id = 3 AND permissao_id = perm.id
        ) THEN
            INSERT INTO t_perfil_permissao (perfil_id, permissao_id, data_hora)
            VALUES (3, perm.id, NOW());
        END IF;
    END LOOP;
END $$;
