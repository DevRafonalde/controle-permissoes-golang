package service

import (
	"context"
	"errors"
	"si-admin/app/helpers"
	"si-admin/app/model/erros"
	"si-admin/app/model/grpc"
	"si-admin/app/model/repositories"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura do serviço de Perfil, contendo repositórios necessários
type PerfilService struct {
	perfilPermissaoRepository repositories.PerfilPermissaoRepository
	permissaoRepository       repositories.PermissaoRepository
	usuarioPerfilRepository   repositories.UsuarioPerfilRepository
	perfilRepository          repositories.PerfilRepository
	usuarioInternoRepository  repositories.UsuarioInternoRepository
	usuarioExternoRepository  repositories.UsuarioExternoRepository
}

// Função para criar uma nova instância de PerfilService com os repositórios necessários
func NewPerfilService(perfilPermissaoRepository repositories.PerfilPermissaoRepository,
	permissaoRepository repositories.PermissaoRepository,
	usuarioPerfilRepository repositories.UsuarioPerfilRepository,
	perfilRepository repositories.PerfilRepository, usuarioInternoRepository repositories.UsuarioInternoRepository, usuarioExternoRepository repositories.UsuarioExternoRepository) *PerfilService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &PerfilService{
		perfilPermissaoRepository: perfilPermissaoRepository,
		permissaoRepository:       permissaoRepository,
		usuarioPerfilRepository:   usuarioPerfilRepository,
		perfilRepository:          perfilRepository,
		usuarioInternoRepository:  usuarioInternoRepository,
		usuarioExternoRepository:  usuarioExternoRepository,
	}
}

// Função para buscar todos os perfis
// Como um perfil pode estar vinculado a várias permissoes,
// Nessa função, os perfis são retornados sem as permissões vinculadas
func (perfilService *PerfilService) FindAllPerfis(context context.Context, req *grpc.RequestAllPaginado) (*grpc.ListaPerfis, erros.ErroStatus) {
	if req.GetTamanhoPagina() == 0 {
		req.TamanhoPagina = 10
	}

	tPerfis, err := perfilService.perfilRepository.FindAll(context, req.GetCursor(), (req.GetTamanhoPagina() + 1))
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum perfil, retorna code NotFound
	if len(tPerfis) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum perfil encontrado"),
		}
	}

	var proximoCursor int32
	temMais := false

	if len(tPerfis) == (int(req.GetTamanhoPagina()) + 1) {
		temMais = true
	}

	var pbPerfis []*grpc.Perfil
	for i, perfil := range tPerfis {
		if i == (int(req.GetTamanhoPagina())) {
			break
		}

		pbPerfis = append(pbPerfis, helpers.TPerfToPb(perfil))
		proximoCursor = perfil.ID
	}

	return &grpc.ListaPerfis{Perfis: pbPerfis, Meta: &grpc.Meta{ProximoCursor: proximoCursor, TamanhoPagina: req.GetTamanhoPagina(), TemMais: temMais}}, erros.ErroStatus{}
}

// Função para buscar um perfil pelo ID
// Diferente da busca por todos, aqui a aplicação retorna o perfil com todas as permissões vinculadas
func (perfilService *PerfilService) FindPerfilById(context context.Context, id int32) (*grpc.PerfilPermissoes, erros.ErroStatus) {
	perfilEncontrado, err := perfilService.perfilRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum perfil, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Perfil não encontrado"),
			}
		}

		return &grpc.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	permissoes, erroPermissoes := perfilService.GetPermissoesVinculadas(context, perfilEncontrado.ID)
	if erroPermissoes.Erro != nil {
		return &grpc.PerfilPermissoes{}, erroPermissoes
	}

	pbPerfilPermissao := grpc.PerfilPermissoes{
		Perfil:     helpers.TPerfToPb(perfilEncontrado),
		Permissoes: permissoes,
	}

	return &pbPerfilPermissao, erros.ErroStatus{}
}

// Função para obter as permissões vinculadas a um perfil, filtrando apenas as ativas
func (perfilService *PerfilService) GetPermissoesVinculadas(context context.Context, id int32) ([]*grpc.Permissao, erros.ErroStatus) {
	perfilPermissoes, errBuscaPermissoes := perfilService.perfilPermissaoRepository.FindByPerfil(context, id)
	if errBuscaPermissoes != nil {
		// Caso não seja encontrado nenhum perfil, retorna code NotFound
		if errBuscaPermissoes.Error() == "no rows in result set" {
			return []*grpc.Permissao{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Perfil não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaPermissoes,
		}
	}

	// Filtra apenas as permissões que estão vinculadas e ativas
	var permissoes []*grpc.Permissao
	for _, perfilPermissao := range perfilPermissoes {
		permissaoEncontrada, err := perfilService.permissaoRepository.FindByID(context, perfilPermissao.PermissaoID)
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		if permissaoEncontrada.Ativo.Bool {
			permissoes = append(permissoes, helpers.TPermissaoToPb(permissaoEncontrada))
		}
	}

	return permissoes, erros.ErroStatus{}
}

// Função para obter os usuários vinculados a um perfil, filtrando apenas os ativos
func (perfilService *PerfilService) GetUsuariosVinculados(context context.Context, id int32) (*grpc.ResponseGetUsuariosVinculados, erros.ErroStatus) {
	perfilUsuarios, errBuscaUsuarios := perfilService.usuarioPerfilRepository.FindByPerfil(context, id)
	if errBuscaUsuarios != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaUsuarios,
		}
	}

	// Filtra apenas os usuários que estão vinculados e ativos
	var usuariosInternos []*grpc.UsuarioInterno
	var usuariosExternos []*grpc.UsuarioExterno
	for _, perfilUsuario := range perfilUsuarios {
		if perfilUsuario.UsuarioInternoID.Int32 != 0 {
			usuarioEncontrado, err := perfilService.usuarioInternoRepository.FindByID(context, perfilUsuario.UsuarioInternoID.Int32)
			if err != nil {
				return nil, erros.ErroStatus{
					Status: codes.Internal,
					Erro:   err,
				}
			}

			if usuarioEncontrado.Ativo.Bool {
				usuariosInternos = append(usuariosInternos, helpers.TUsuarioInternoToPb(usuarioEncontrado))
			}
		}

		if perfilUsuario.UsuarioExternoID.Int32 != 0 {
			usuarioEncontrado, err := perfilService.usuarioExternoRepository.FindByID(context, perfilUsuario.UsuarioExternoID.Int32)
			if err != nil {
				return nil, erros.ErroStatus{
					Status: codes.Internal,
					Erro:   err,
				}
			}

			if usuarioEncontrado.Ativo.Bool {
				usuariosExternos = append(usuariosExternos, helpers.TUsuarioExternoToPb(usuarioEncontrado))
			}
		}
	}

	return &grpc.ResponseGetUsuariosVinculados{UsuariosInternos: usuariosInternos, UsuariosExternos: usuariosExternos}, erros.ErroStatus{}
}

// Função para clonar um perfil, mantendo as permissões, mas retornando um novo perfil vazio
func (perfilService *PerfilService) ClonePerfil(context context.Context, id int32) (*grpc.PerfilPermissoes, erros.ErroStatus) {
	perfil, err := perfilService.FindPerfilById(context, id)
	if err.Erro != nil {
		return &grpc.PerfilPermissoes{}, err
	}

	// Define o perfil clonado como vazio
	perfil.Perfil = &grpc.Perfil{}

	return perfil, erros.ErroStatus{}
}

// Função para criar um novo perfil com permissões
func (perfilService *PerfilService) CreatePerfil(context context.Context, perfilPermissaoRecebido *grpc.PerfilPermissoes) (*grpc.PerfilPermissoes, erros.ErroStatus) {
	// Busca um perfil pelo nome enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	perfisEncontrados, erroBuscaPreExistente := perfilService.perfilRepository.FindByNome(context, perfilPermissaoRecebido.Perfil.GetNome(), 0, 9999)
	if erroBuscaPreExistente != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroBuscaPreExistente,
		}
	}

	if len(perfisEncontrados) != 0 {
		for _, perfil := range perfisEncontrados {
			if strings.Compare(perfil.Nome, perfilPermissaoRecebido.Perfil.GetNome()) == 0 {
				return nil, erros.ErroStatus{
					Status: codes.AlreadyExists,
					Erro:   errors.New("Já existe perfil com o nome enviado"),
				}
			}
		}
	}

	// Cria o objeto CreatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilCreate := db.CreatePerfilParams{
		Nome:      perfilPermissaoRecebido.Perfil.Nome,
		Descricao: perfilPermissaoRecebido.Perfil.Descricao,
	}

	// Cria o perfil no repositório
	perfil, err := perfilService.perfilRepository.Create(context, perfilCreate)
	perfilPermissaoRecebido.Perfil = helpers.TPerfToPb(perfil)
	if err != nil {
		return &grpc.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Salva a relação entre o perfil e cada permissão
	for i, permissao := range perfilPermissaoRecebido.Permissoes {
		tPermissao, err := perfilService.permissaoRepository.FindByID(context, permissao.Id)
		if err != nil {
			return &grpc.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo da permissão na lista de permissões do objeto recebido
		// Como esse é o objeto retornado, as permissões retornam completas e prontas para serem usadas
		perfilPermissaoRecebido.Permissoes[i] = helpers.TPermissaoToPb(tPermissao)

		// Cria o objeto CreatePerfilPermissaoParams gerado pelo sqlc para gravação no banco de dados
		perfilPermissao := db.CreatePerfilPermissaoParams{
			PerfilID:    perfil.ID,
			PermissaoID: permissao.Id,
			DataHora:    pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		perfilService.perfilPermissaoRepository.Create(context, perfilPermissao)
	}

	return perfilPermissaoRecebido, erros.ErroStatus{}
}

// Função para atualizar um perfil e suas permissões
func (perfilService *PerfilService) UpdatePerfil(context context.Context, perfilPermissao *grpc.PerfilPermissoes, perfilPermissaoAntigo *grpc.PerfilPermissoes) (*grpc.PerfilPermissoes, erros.ErroStatus) {
	perfilRecebido := perfilPermissao.Perfil

	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if perfilPermissaoAntigo.GetPerfil().GetNome() != perfilRecebido.GetNome() {
		perfisEncontrados, erroBuscaPreExistente := perfilService.perfilRepository.FindByNome(context, perfilRecebido.GetNome(), 0, 9999)
		if erroBuscaPreExistente != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   erroBuscaPreExistente,
			}
		}

		if len(perfisEncontrados) != 0 {
			for _, perfil := range perfisEncontrados {
				if strings.Compare(perfil.Nome, perfilRecebido.GetNome()) == 0 {
					return nil, erros.ErroStatus{
						Status: codes.AlreadyExists,
						Erro:   errors.New("Já existe perfil com o nome enviado"),
					}
				}
			}
		}
	}

	// Cria o objeto UpdatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilUpdate := db.UpdatePerfilParams{
		Nome:      perfilRecebido.Nome,
		Descricao: perfilRecebido.Descricao,
		ID:        perfilRecebido.Id,
	}

	// Salva o perfil atualizado no repositório
	perfilAtualizado, err := perfilService.perfilRepository.Update(context, perfilUpdate)
	if err != nil {
		return &grpc.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Deleta todas as relações antigas entre o perfil e suas permissões
	registrosExistentes, errBuscaPermissoes := perfilService.perfilPermissaoRepository.FindByPerfil(context, perfilAtualizado.ID)
	if errBuscaPermissoes != nil {
		return &grpc.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	for i := 0; i < len(registrosExistentes); i++ {
		erroDelete := perfilService.perfilPermissaoRepository.Delete(context, registrosExistentes[i].ID)
		if erroDelete != nil {
			return &grpc.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   erroDelete,
			}
		}
	}

	// Cria novas relações entre o perfil atualizado e as novas permissões
	permissoes := perfilPermissao.Permissoes

	for i, permissao := range permissoes {
		tPermissao, err := perfilService.permissaoRepository.FindByID(context, permissao.Id)
		if err != nil {
			if err.Error() == "no rows in result set" {
				return nil, erros.ErroStatus{
					Status: codes.NotFound,
					Erro:   errors.New("Permissão não encontrada"),
				}
			}

			return &grpc.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo da permissão na lista de permissões do objeto recebido
		// Como esse é o objeto retornado, as permissões retornam completas e prontas para serem usadas
		permissoes[i] = helpers.TPermissaoToPb(tPermissao)

		// Cria o objeto CreatePerfilPermissaoParams gerado pelo sqlc para gravação no banco de dados
		perfilPermissao := db.CreatePerfilPermissaoParams{
			PerfilID:    perfilAtualizado.ID,
			PermissaoID: permissao.Id,
			DataHora:    pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		perfilService.perfilPermissaoRepository.Create(context, perfilPermissao)
	}

	perfilPermissaoRetorno := grpc.PerfilPermissoes{
		Perfil:     helpers.TPerfToPb(perfilAtualizado),
		Permissoes: permissoes,
	}

	return &perfilPermissaoRetorno, erros.ErroStatus{}
}

// Função para desativar um perfil pelo ID
func (perfilService *PerfilService) DesativarPerfilById(context context.Context, req *grpc.PerfilDelete) erros.ErroStatus {
	perfilDeletado, erroService := perfilService.FindPerfilById(context, req.GetIdPerfilDeletado())
	if erroService.Erro != nil {
		return erroService
	}

	// Cria o objeto UpdatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilUpdate := db.UpdatePerfilParams{
		Nome:      perfilDeletado.GetPerfil().GetNome(),
		Descricao: perfilDeletado.GetPerfil().GetDescricao(),
		ID:        perfilDeletado.GetPerfil().GetId(),
	}

	// Salva o perfil atualizado no repositório
	_, err := perfilService.perfilRepository.Update(context, perfilUpdate)
	if err != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	usuarioPerfis, err := perfilService.usuarioPerfilRepository.FindByPerfil(context, req.GetIdPerfilDeletado())
	if err != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	for _, usuarioPerfil := range usuarioPerfis {
		_, err := perfilService.usuarioPerfilRepository.Update(context, db.UpdateUsuarioPerfilParams{
			UsuarioInternoID: usuarioPerfil.UsuarioInternoID,
			UsuarioExternoID: usuarioPerfil.UsuarioExternoID,
			PerfilID:         req.GetIdPerfilNovo(),
			DataHora:         pgtype.Timestamp{Time: time.Now(), Valid: true},
			ID:               usuarioPerfil.ID,
		})

		if err != nil {
			return erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	}

	return erros.ErroStatus{}
}

// Função para ativar um perfil pelo ID
func (perfilService *PerfilService) RestaurarPerfilById(context context.Context, perfil *grpc.PerfilPermissoes) erros.ErroStatus {
	// Cria o objeto UpdatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilUpdate := db.UpdatePerfilParams{
		Nome:      perfil.Perfil.Nome,
		Descricao: perfil.Perfil.Descricao,
		ID:        perfil.Perfil.Id,
	}

	// Salva o perfil atualizado no repositório
	_, errAtt := perfilService.perfilRepository.Update(context, perfilUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}
