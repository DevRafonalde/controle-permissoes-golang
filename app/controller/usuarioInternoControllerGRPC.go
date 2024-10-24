package controller

import (
	"context"
	"errors"
	"si-admin/app/configuration/logger"
	"si-admin/app/controller/middlewares"
	"si-admin/app/service"
	"strconv"
	"time"

	pb "si-admin/app/model/grpc" // Importa o pacote gerado pelos arquivos .proto

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implementação do servidor
type UsuariosInternosServer struct {
	pb.UnimplementedUsuariosInternosServer
	usuarioInternoService *service.UsuarioInternoService
	permissaoMiddleware   *middlewares.PermissoesMiddleware
	chaveSecreta          []byte
}

func NewUsuariosInternosServer(usuarioInternoService *service.UsuarioInternoService, permissaoMiddleware *middlewares.PermissoesMiddleware, chaveSecreta []byte) *UsuariosInternosServer {
	return &UsuariosInternosServer{
		usuarioInternoService: usuarioInternoService,
		permissaoMiddleware:   permissaoMiddleware,
		chaveSecreta:          chaveSecreta,
	}
}

func (usuarioInternoServer *UsuariosInternosServer) mustEmbedUnimplementedUsuariosInternosServer() {}

// Função para buscar por todos os usuários
func (usuarioInternoServer *UsuariosInternosServer) FindAllUsuarios(context context.Context, req *pb.RequestAllPaginado) (*pb.ListaUsuariosInternos, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-usuarios-internos")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaUsuariosInternos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarios, erroService := usuarioInternoServer.usuarioInternoService.FindAllUsuariosInternos(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.ListaUsuariosInternos{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados os "+strconv.Itoa(int(req.GetTamanhoPagina()))+" primeiros registros de usuários a partir do ID"+strconv.Itoa(int(req.GetCursor())),
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
	)

	return usuarios, nil
}

// Função para buscar por um usuário pelo Id
func (usuarioInternoServer *UsuariosInternosServer) FindUsuarioById(context context.Context, req *pb.RequestId) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-interno-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.UsuarioPerfis{}, status.Errorf(codes.InvalidArgument, "Id não enviado")
	}

	usuario, erroService := usuarioInternoServer.usuarioInternoService.FindUsuarioInternoById(context, id) // Busca o usuário pelo Id
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	pbPerfis = append(pbPerfis, usuario.Perfis...)

	usuarioRetorno := &pb.UsuarioPerfis{UsuarioInterno: usuario.UsuarioInterno, Perfis: pbPerfis}

	logger.Logger.Info("Buscado um usuário pelo Id",
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
		zap.Any("usuarioBuscado", usuarioRetorno),
	)

	return usuarioRetorno, nil
}

// Função para buscar por todos os perfis vinculados àquele usuário
func (usuarioInternoServer *UsuariosInternosServer) GetPerfisVinculados(context context.Context, req *pb.RequestId) (*pb.ResponsePerfisVinculados, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-perfis-vinculados")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponsePerfisVinculados{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponsePerfisVinculados{}, status.Errorf(codes.InvalidArgument, "Id não enviado")
	}

	perfis, erroService := usuarioInternoServer.usuarioInternoService.GetPerfisVinculados(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.ResponsePerfisVinculados{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	logger.Logger.Info("Buscado os perfis vinculados a um usuário pelo Id do usuário",
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
	)

	return &pb.ResponsePerfisVinculados{Perfis: pbPerfis}, nil
}

// Função para criar um novo usuário
func (usuarioInternoServer *UsuariosInternosServer) CreateUsuario(context context.Context, req *pb.UsuarioPerfis) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-usuario-interno")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioCriado, erroService := usuarioInternoServer.usuarioInternoService.CreateUsuarioInterno(context, req) // Cria o usuário
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo usuário",
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
		zap.Any("usuarioCriado", usuarioCriado),
	)

	return usuarioCriado, nil
}

// Função para clonar um usuário existente
func (usuarioInternoServer *UsuariosInternosServer) CloneUsuario(context context.Context, req *pb.RequestId) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-clone-usuario-interno")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.UsuarioPerfis{}, status.Errorf(codes.InvalidArgument, "Id não enviado")
	}

	usuario, erroService := usuarioInternoServer.usuarioInternoService.CloneUsuarioInterno(context, id) // Clona o usuário
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	pbPerfis = append(pbPerfis, usuario.Perfis...)

	logger.Logger.Info("Clonado um usuário existente",
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
		zap.Any("usuarioClonado", usuario),
	)

	return &pb.UsuarioPerfis{UsuarioInterno: usuario.UsuarioInterno, Perfis: pbPerfis}, nil
}

// Função para atualizar um usuário já existente
func (usuarioInternoServer *UsuariosInternosServer) UpdateUsuario(context context.Context, req *pb.UsuarioPerfis) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-usuario-interno")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioAntigo, erroService := usuarioInternoServer.usuarioInternoService.FindUsuarioInternoById(context, req.GetUsuarioInterno().GetId()) // Busca o usuário pelo Id
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	usuarioNovo, erroService := usuarioInternoServer.usuarioInternoService.UpdateUsuarioInterno(context, req, usuarioAntigo) // Atualiza o usuário
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	pbPerfis = append(pbPerfis, usuarioNovo.Perfis...)

	logger.Logger.Info("Atualizado um usuário existente",
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
		zap.Any("usuarioAntigo", usuarioAntigo),
		zap.Any("usuarioAtualizado", usuarioNovo),
	)

	return &pb.UsuarioPerfis{UsuarioInterno: usuarioNovo.UsuarioInterno, Perfis: pbPerfis}, nil
}

// Função para o administrador do sistema alterar a senha de qualquer usuário baseado no Id
func (usuarioInternoServer *UsuariosInternosServer) AlterarSenhaAdmin(context context.Context, req *pb.RequestAlterarSenhaAdmin) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-alterar-senha-admin")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id não enviado")
	}

	usuarioASerAlterado, erroService := usuarioInternoServer.usuarioInternoService.FindUsuarioInternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = usuarioInternoServer.usuarioInternoService.AtualizarSenha(context, usuarioASerAlterado.GetUsuarioInterno().GetEmail(), req.GetSenhaNova())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Alterada a senha de um usuário existente por meio de permissões de admin",
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
		zap.Any("usuarioModificado", usuarioASerAlterado),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para o usuário alterar a própria senha
func (usuarioInternoServer *UsuariosInternosServer) AlterarSenhaUsuario(context context.Context, req *pb.RequestAlterarSenhaUsuario) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-alterar-propria-senha")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id não enviado")
	}

	usuarioASerAlterado, erroService := usuarioInternoServer.usuarioInternoService.FindUsuarioInternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	login := pb.LoginUsuario{
		Email: usuarioASerAlterado.UsuarioInterno.Email,
		Senha: req.SenhaAntiga,
	}

	usuarioLogado, erroService := usuarioInternoServer.usuarioInternoService.Login(context, &login)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = usuarioInternoServer.usuarioInternoService.AtualizarSenha(context, usuarioLogado.UsuarioInterno.Email, req.SenhaNova)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioInterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Alterada a senha de um usuário existente por meios próprios",
		zap.Any("usuario", usuarioSolicitante.UsuarioInterno),
		zap.Any("usuarioModificado", usuarioASerAlterado),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para restaurar um usuário existente
func (usuarioInternoServer *UsuariosInternosServer) RestaurarUsuario(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-restaurar-usuario-interno")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	usuarioExterno, erroService := usuarioInternoServer.usuarioInternoService.FindUsuarioInternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	restaurado, erroService := usuarioInternoServer.usuarioInternoService.RestaurarUsuarioInterno(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !restaurado {
		logger.Logger.Error("Não existe usuário externo com o Id enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe usuário externo com o Id enviado")
	}

	logger.Logger.Info("Restaurado um usuário externo",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioExterno", usuarioExterno),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para desativar um usuário existente
func (usuarioInternoServer *UsuariosInternosServer) DesativarUsuario(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioInternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-desativar-usuario-interno")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	usuarioExterno, erroService := usuarioInternoServer.usuarioInternoService.FindUsuarioInternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	desativado, erroService := usuarioInternoServer.usuarioInternoService.DesativarUsuarioInterno(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !desativado {
		logger.Logger.Error("Não existe usuário externo com o Id enviado", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe usuário externo com o Id enviado")
	}

	logger.Logger.Info("Desativado um usuário externo",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioExterno", usuarioExterno),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para realizar o login na aplicação
func (usuarioInternoServer *UsuariosInternosServer) Login(context context.Context, req *pb.LoginUsuario) (*pb.RetornoLoginUsuario, error) {
	usuarioLogado, erroService := usuarioInternoServer.usuarioInternoService.Login(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Tentativa de login na API", zap.NamedError("err", erroService.Erro), zap.Any("email", req.GetEmail()))
		return &pb.RetornoLoginUsuario{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	tokenString, err := usuarioInternoServer.createToken(context, usuarioLogado)
	if err != nil {
		logger.Logger.Error(err.Error(), zap.NamedError("err", err), zap.Any("usuario", usuarioLogado.UsuarioInterno))
		return &pb.RetornoLoginUsuario{}, status.Errorf(codes.Internal, err.Error())
	}

	logger.Logger.Info("Feito login na API",
		zap.Any("usuario", usuarioLogado.UsuarioInterno),
	)

	return &pb.RetornoLoginUsuario{
		Id:    usuarioLogado.GetUsuarioInterno().GetId(),
		Nome:  usuarioLogado.GetUsuarioInterno().GetNome(),
		Email: usuarioLogado.GetUsuarioInterno().Email,
		Token: tokenString,
	}, nil
}

// Função para envio do token de reset de senha
func (usuarioInternoServer *UsuariosInternosServer) TokenResetSenha(context context.Context, req *pb.EmailReset) (*pb.ResponseTokenResetSenha, error) {
	token, err := usuarioInternoServer.createTokenReset(req)
	if err != nil {
		logger.Logger.Error(err.Error(), zap.NamedError("err", err), zap.Any("email", req.GetEmail()))
		return &pb.ResponseTokenResetSenha{}, status.Errorf(codes.Internal, err.Error())
	}

	erroService := usuarioInternoServer.usuarioInternoService.TokenResetSenha(context, token, req.GetEmail())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("email", req.GetEmail()))
		return &pb.ResponseTokenResetSenha{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Solicitado um token para reset/esquecimento de senha",
		zap.Any("email", req.GetEmail()),
	)

	return &pb.ResponseTokenResetSenha{Token: token}, nil
}

// Função para validar o token de reset de senha e alterar a senha
func (usuarioInternoServer *UsuariosInternosServer) ResetSenha(context context.Context, req *pb.ResetSenhaUsuario) (*pb.ResponseBool, error) {
	email, err := usuarioInternoServer.validarToken(req.Token)
	if err != nil {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.Unauthenticated, "Token inválido")
	}

	erroService := usuarioInternoServer.usuarioInternoService.AtualizarSenha(context, email, req.GetSenhaNova())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("email", email))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Resetada e alterada a senha de usuário",
		zap.Any("email", email),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// createToken cria um token JWT para um usuário autenticado.
func (usuarioInternoServer *UsuariosInternosServer) createToken(context context.Context, usuarioLogado *pb.UsuarioPerfis) (string, error) {
	claims := middlewares.CustomClaims{
		IdUsuario:      usuarioLogado.UsuarioInterno.Id,
		UsuarioInterno: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24), // Define a expiração do token para 24 horas
			},
		},
	}

	logger.Logger.Info("Solicitada a criação de um novo token de acesso",
		zap.Any("usuario", usuarioLogado.UsuarioInterno),
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Cria o token JWT com as claims

	tokenString, err := token.SignedString(usuarioInternoServer.chaveSecreta) // Assina o token com a chave secreta
	if err != nil {
		logger.Logger.Error("Falha na criação do token de acesso", zap.NamedError("err", err), zap.Any("usuario", usuarioLogado.UsuarioInterno))
		return "", err // Retorna erro se falhar
	}

	_, erroService := usuarioInternoServer.usuarioInternoService.SenhaAtualizada(context, usuarioLogado.GetUsuarioInterno().GetId())
	if erroService.Erro != nil {
		logger.Logger.Error("Falha na criação do token de acesso", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioLogado.UsuarioInterno))
		return "", erroService.Erro
	}

	logger.Logger.Info("Criado um novo token de acesso",
		zap.Any("usuario", usuarioLogado.UsuarioInterno),
	)

	return tokenString, nil // Retorna o token JWT gerado
}

// createToken cria um token JWT para um usuário autenticado para fins de reset da senha do mesmo.
func (usuarioInternoServer *UsuariosInternosServer) createTokenReset(emailReset *pb.EmailReset) (string, error) {
	claims := CustomClaimsResetSenha{
		Email:          emailReset.Email,
		UsuarioInterno: true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * 5), // Define a expiração do token para 5 minutos
			},
		},
	}

	logger.Logger.Info("Solicitada a criação de um novo token de reset de senha",
		zap.Any("email", emailReset.GetEmail()),
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Cria o token JWT com as claims

	tokenString, err := token.SignedString(usuarioInternoServer.chaveSecreta) // Assina o token com a chave secreta
	if err != nil {
		logger.Logger.Error("Falha na criação do token de reset de senha", zap.NamedError("err", err), zap.Any("email", emailReset.GetEmail()))
		return "", err // Retorna erro se falhar
	}

	logger.Logger.Info("Criado um novo token de reset de senha",
		zap.Any("email", emailReset.GetEmail()),
	)

	return tokenString, nil // Retorna o token JWT gerado
}

// Struct para capturar as informações do token JWT de reset de senha
type CustomClaimsResetSenha struct {
	Email                string
	UsuarioInterno       bool
	jwt.RegisteredClaims // Struct que contém os claims padrão do JWT
}

// Função para verificar a validade de um token JWT
func (usuarioInternoServer *UsuariosInternosServer) validarToken(tokenString string) (string, error) {
	// Analisa o token JWT usando a chave secreta
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsResetSenha{}, func(token *jwt.Token) (interface{}, error) {
		return usuarioInternoServer.chaveSecreta, nil
	})
	if err != nil {
		logger.Logger.Error("Falha na validação de um token de reset de senha", zap.NamedError("err", err))
		return "", err // Retorna erro se o token não puder ser analisado
	}

	if !token.Valid {
		logger.Logger.Error("Enviado um token de reset de senha inválido", zap.NamedError("err", err))
		return "", errors.New("Token inválido") // Retorna erro se o token não for válido
	}

	// Extrai as claims personalizadas do token
	props, ok := token.Claims.(*CustomClaimsResetSenha)
	if !ok {
		logger.Logger.Error("Falha na extração das propriedades de um token de reset de senha", zap.NamedError("err", err))
		return "", errors.New("Propriedades inválidas do token") // Retorna erro se as claims não puderem ser extraídas
	}

	email := props.Email // Obtém o email do usuário a partir das claims

	logger.Logger.Info("Validado um token de reset de senha",
		zap.Any("email", email),
	)

	return email, nil // Retorna o email do usuário
}
