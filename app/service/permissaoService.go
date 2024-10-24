package service

import (
	"context"
	"errors"
	"si-admin/app/helpers"
	"si-admin/app/model/erros"
	pb "si-admin/app/model/grpc"
	"si-admin/app/model/repositories"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
)

// Estrutura do serviço de Permissão, contendo repositórios necessários
type PermissaoService struct {
	perfilPermissaoRepository repositories.PerfilPermissaoRepository
	permissaoRepository       repositories.PermissaoRepository
}

// Função para criar uma nova instância de PermissaoService com os repositórios necessários
func NewPermissaoService(perfilPermissaoRepository repositories.PerfilPermissaoRepository, permissaoRepository repositories.PermissaoRepository) *PermissaoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &PermissaoService{
		perfilPermissaoRepository: perfilPermissaoRepository,
		permissaoRepository:       permissaoRepository,
	}
}

// Função para buscar todas as permissões
func (permissaoService *PermissaoService) FindAllPermissoes(context context.Context, req *pb.RequestAllPaginado) (*pb.ListaPermissoes, erros.ErroStatus) {
	if req.GetTamanhoPagina() == 0 {
		req.TamanhoPagina = 10
	}

	tPermissoes, err := permissaoService.permissaoRepository.FindAll(context, req.GetCursor(), (req.GetTamanhoPagina() + 1))
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhuma permissão, retorna code NotFound
	if len(tPermissoes) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma permissão encontrado"),
		}
	}

	var proximoCursor int32
	temMais := false

	if len(tPermissoes) == (int(req.GetTamanhoPagina()) + 1) {
		temMais = true
	}

	var pbPermissoes []*pb.Permissao
	for i, permissao := range tPermissoes {
		if i == (int(req.GetTamanhoPagina())) {
			break
		}

		pbPermissoes = append(pbPermissoes, helpers.TPermissaoToPb(permissao))
		proximoCursor = permissao.ID
	}

	return &pb.ListaPermissoes{Permissoes: pbPermissoes, Meta: &pb.Meta{ProximoCursor: proximoCursor, TamanhoPagina: req.GetTamanhoPagina(), TemMais: temMais}}, erros.ErroStatus{}
}

// Função para buscar uma permissão pelo Id
func (permissaoService *PermissaoService) FindPermissaoById(context context.Context, id int32) (*pb.Permissao, erros.ErroStatus) {
	permissao, err := permissaoService.permissaoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhuma permissão, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Permissao{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Permissão não encontrado"),
			}
		}

		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TPermissaoToPb(permissao), erros.ErroStatus{}
}

// Função para buscar uma permissão pelo nome
func (permissaoService *PermissaoService) FindPermissaoByNome(context context.Context, req *pb.RequestNome) ([]*pb.Permissao, erros.ErroStatus) {
	if req.GetTamanhoPagina() == 0 {
		req.TamanhoPagina = 10
	}

	tPermissoes, err := permissaoService.permissaoRepository.FindByNome(context, req.GetNome(), req.GetCursor(), (req.GetTamanhoPagina() + 1))
	if err != nil {
		return []*pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhuma permissão, retorna code NotFound
	if len(tPermissoes) == 0 {
		return []*pb.Permissao{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma permissão encontrado"),
		}
	}

	var pbPermissoes []*pb.Permissao
	for _, permissao := range tPermissoes {
		pbPermissoes = append(pbPermissoes, helpers.TPermissaoToPb(permissao))
	}

	return pbPermissoes, erros.ErroStatus{}
}

// Função para criar uma nova permissão
func (permissaoService *PermissaoService) CreatePermissao(context context.Context, permissaoRecebida *pb.Permissao) (*pb.Permissao, erros.ErroStatus) {
	// Busca uma permissão pelo nome enviado para verificar a prévia existência dela
	// Em caso positivo, retorna code AlreadyExists
	permissoesEncontradas, err := permissaoService.permissaoRepository.FindByNome(context, permissaoRecebida.GetNome(), 0, 9999)
	if err != nil {
		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	if len(permissoesEncontradas) != 0 {
		for _, permissao := range permissoesEncontradas {
			if strings.Compare(permissao.Nome, permissaoRecebida.GetNome()) == 0 {
				return nil, erros.ErroStatus{
					Status: codes.AlreadyExists,
					Erro:   errors.New("Já existe permissão com o nome enviado"),
				}
			}
		}
	}

	// Cria o objeto CreatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoCreate := db.CreatePermissaoParams{
		Nome:      permissaoRecebida.GetNome(),
		Descricao: permissaoRecebida.GetDescricao(),
	}

	// Cria a permissão no repositório
	permissaoCriada, err := permissaoService.permissaoRepository.Create(context, permissaoCreate)
	if err != nil {
		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TPermissaoToPb(permissaoCriada), erros.ErroStatus{}
}

// Função para atualizar uma permissão existente
func (permissaoService *PermissaoService) UpdatePermissao(context context.Context, permissaoRecebida *pb.Permissao, permissaoBanco *pb.Permissao) (*pb.Permissao, erros.ErroStatus) {
	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if permissaoBanco.GetNome() != permissaoRecebida.GetNome() {
		permissoesEncontradas, err := permissaoService.permissaoRepository.FindByNome(context, permissaoRecebida.GetNome(), 0, 9999)
		if err != nil {
			return &pb.Permissao{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		if len(permissoesEncontradas) != 0 {
			for _, permissao := range permissoesEncontradas {
				if strings.Compare(permissao.Nome, permissaoRecebida.GetNome()) == 0 {
					return nil, erros.ErroStatus{
						Status: codes.AlreadyExists,
						Erro:   errors.New("Já existe permissão com o nome enviado"),
					}
				}
			}
		}
	}

	// Cria o objeto UpdatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoUpdate := db.UpdatePermissaoParams{
		Nome:      permissaoRecebida.GetNome(),
		Descricao: permissaoRecebida.GetDescricao(),
		ID:        permissaoBanco.GetId(),
	}

	// Salva a permissão atualizada no repositório
	permissaoAtualizada, erroUpdate := permissaoService.permissaoRepository.Update(context, permissaoUpdate)
	if erroUpdate != nil {
		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroUpdate,
		}
	}

	return helpers.TPermissaoToPb(permissaoAtualizada), erros.ErroStatus{}
}

// Função para desativar uma permissão pelo Id
func (permissaoService *PermissaoService) DesativarPermissaoById(context context.Context, permissao *pb.Permissao) erros.ErroStatus {
	// Cria o objeto UpdatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoUpdate := db.UpdatePermissaoParams{
		Nome:      permissao.Nome,
		Descricao: permissao.Descricao,
		ID:        permissao.Id,
	}

	// Salva a permissão atualizada no repositório
	_, errAtt := permissaoService.permissaoRepository.Update(context, permissaoUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}

// Função para ativar uma permissão pelo Id
func (permissaoService *PermissaoService) RestaurarPermissaoById(context context.Context, permissao *pb.Permissao) erros.ErroStatus {
	// Cria o objeto UpdatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoUpdate := db.UpdatePermissaoParams{
		Nome:      permissao.Nome,
		Descricao: permissao.Descricao,
		ID:        permissao.Id,
	}

	// Salva a permissão atualizada no repositório
	_, errAtt := permissaoService.permissaoRepository.Update(context, permissaoUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}
