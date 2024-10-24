package controller

import (
	"context"
	"si-admin/app/configuration/logger"
	"si-admin/app/controller/middlewares"
	"si-admin/app/service"
	"strconv"

	pb "si-admin/app/model/grpc" // Importa o pacote gerado pelos arquivos .proto

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implementação do servidor
type PermissoesServer struct {
	pb.UnimplementedPermissoesServer
	permissaoService    *service.PermissaoService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewPermissoesServer(permissaoService *service.PermissaoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *PermissoesServer {
	return &PermissoesServer{
		permissaoService:    permissaoService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (permissoesServer *PermissoesServer) mustEmbedUnimplementedPermissoesServer() {}

// Função para buscar por todas as permissões
func (permissoesServer *PermissoesServer) FindAllPermissoes(context context.Context, req *pb.RequestAllPaginado) (*pb.ListaPermissoes, error) {
	usuarioSolicitante, retornoMiddleware := permissoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-permissoes")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaPermissoes{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	permissoes, erroService := permissoesServer.permissaoService.FindAllPermissoes(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados os "+strconv.Itoa(int(req.GetTamanhoPagina()))+" primeiros registros de permissões a partir do ID"+strconv.Itoa(int(req.GetCursor())),
		zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()),
	)

	return permissoes, nil
}

// Função para buscar por uma permissão pelo Id
func (permissoesServer *PermissoesServer) FindPermissaoById(context context.Context, req *pb.RequestId) (*pb.Permissao, error) {
	usuarioSolicitante, retornoMiddleware := permissoesServer.permissaoMiddleware.PermissaoMiddleware(context, "get-permissao-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Permissao{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	permissao, erroService := permissoesServer.permissaoService.FindPermissaoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada uma permissão por Id",
		zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()),
		zap.Any("permissao", permissao),
	)

	return permissao, nil
}

// Função para criar uma nova permissão
func (permissoesServer *PermissoesServer) CreatePermissao(context context.Context, req *pb.Permissao) (*pb.Permissao, error) {
	usuarioSolicitante, retornoMiddleware := permissoesServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-permissao")
	if retornoMiddleware.Erro != nil {
		return &pb.Permissao{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	permissao, erroService := permissoesServer.permissaoService.CreatePermissao(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criada uma nova permissão",
		zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()),
		zap.Any("permissao", permissao),
	)

	return permissao, nil
}

// Função para atualizar uma permissão já existente no banco
func (permissoesServer *PermissoesServer) UpdatePermissao(context context.Context, req *pb.Permissao) (*pb.Permissao, error) {
	usuarioSolicitante, retornoMiddleware := permissoesServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-permissao")
	if retornoMiddleware.Erro != nil {
		return &pb.Permissao{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	permissaoAntiga, erroService := permissoesServer.permissaoService.FindPermissaoById(context, req.GetId())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	permissao, erroService := permissoesServer.permissaoService.UpdatePermissao(context, req, permissaoAntiga)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizada uma permissão existente",
		zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()),
		zap.Any("permissaoAntiga", permissaoAntiga),
		zap.Any("permissao", permissao),
	)

	return permissao, nil
}

// Função para restaurar uma permissão existente
func (permissoesServer *PermissoesServer) RestaurarPermissao(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := permissoesServer.permissaoMiddleware.PermissaoMiddleware(context, "put-restaurar-permissao")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	permissao, erroService := permissoesServer.permissaoService.FindPermissaoById(context, req.GetId())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = permissoesServer.permissaoService.RestaurarPermissaoById(context, permissao)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Restaurada uma permissão existente",
		zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()),
		zap.Any("permissao", permissao),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para desativar uma permissão existente
func (permissoesServer *PermissoesServer) DesativarPermissao(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := permissoesServer.permissaoMiddleware.PermissaoMiddleware(context, "put-desativar-permissao")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	permissao, erroService := permissoesServer.permissaoService.FindPermissaoById(context, req.GetId())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = permissoesServer.permissaoService.DesativarPermissaoById(context, permissao)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Desativada uma permissão existente",
		zap.Any("usuarioExterno", usuarioSolicitante.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioSolicitante.GetUsuarioInterno()),
		zap.Any("permissao", permissao),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
