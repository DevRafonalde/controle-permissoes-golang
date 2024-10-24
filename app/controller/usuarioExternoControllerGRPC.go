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
type UsuariosExternosServer struct {
	pb.UnimplementedUsuariosExternosServer
	usuarioExternoService *service.UsuarioExternoService
	permissaoMiddleware   *middlewares.PermissoesMiddleware
	chaveSecreta          []byte
}

func NewUsuarioExternoServer(clienteService *service.UsuarioExternoService, permissaoMiddleware *middlewares.PermissoesMiddleware, chaveSecreta []byte) *UsuariosExternosServer {
	return &UsuariosExternosServer{
		usuarioExternoService: clienteService,
		permissaoMiddleware:   permissaoMiddleware,
		chaveSecreta:          chaveSecreta,
	}
}

func (usuarioExternoServer *UsuariosExternosServer) mustEmbedUnimplementedUsuariosExternosServer() {}

// Função para buscar por todos os usuários externos
func (usuarioExternoServer *UsuariosExternosServer) FindAllUsuariosExternos(context context.Context, req *pb.RequestAllPaginado) (*pb.ListaUsuariosExternos, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-usuarios-externos")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaUsuariosExternos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuariosExternos, erroService := usuarioExternoServer.usuarioExternoService.FindAllUsuariosExternos(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar todos os usuários externos "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados os "+strconv.Itoa(int(req.GetTamanhoPagina()))+" primeiros registros de usuários externos a partir do ID"+strconv.Itoa(int(req.GetCursor())),
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
	)

	return usuariosExternos, nil
}

// Função para buscar por um usuário externo pelo Id
func (usuarioExternoServer *UsuariosExternosServer) FindUsuarioExternoByID(context context.Context, req *pb.RequestId) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-externo-by-id")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	usuarioExterno, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar usuário externo pelo Id "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um usuário externo por Id",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioExterno", usuarioExterno),
	)

	return usuarioExterno, nil
}

// Função para buscar por um usuário externo pelo Id externo
func (usuarioExternoServer *UsuariosExternosServer) FindUsuarioExternoByIDExterno(context context.Context, req *pb.RequestId) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-externo-by-id-externo")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	usuarioExterno, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoByIdExterno(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar usuário externo pelo Id externo "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um usuário externo pelo Id externo",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioExterno", usuarioExterno),
	)

	return usuarioExterno, nil
}

// Função para buscar por um usuário externo pelo nome
func (usuarioExternoServer *UsuariosExternosServer) FindUsuarioExternoByNome(context context.Context, req *pb.RequestNome) (*pb.ListaUsuariosExternos, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-externo-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaUsuariosExternos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuariosExternos, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoByNome(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar o usuário externo pelo nome "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado um usuário externo pelo nome",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuariosExternos", usuariosExternos),
	)

	return usuariosExternos, nil
}

// Função para buscar por um usuário externo pelo documento
func (usuarioExternoServer *UsuariosExternosServer) FindUsuarioExternoByDocumento(context context.Context, req *pb.RequestDocumento) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-externo-by-documento")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioExterno, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoByDocumento(context, req.GetDocumento())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar os usuários externos pelo documento "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados os usuários externos pelo documento",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("documento", req.GetDocumento()),
	)

	return usuarioExterno, nil
}

// Função para buscar por um usuário externo pelo documento
func (usuarioExternoServer *UsuariosExternosServer) FindUsuarioExternoByEmail(context context.Context, req *pb.RequestEmail) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-externo-by-email")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioExterno, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoByEmail(context, req.GetEmail())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar os usuários externos pelo e-mail "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados os usuários externos pelo documento",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("email", req.GetEmail()),
	)

	return usuarioExterno, nil
}

// Função para buscar por um usuário externo pelo código de reserva
func (usuarioExternoServer *UsuariosExternosServer) FindUsuarioExternoByCodReserva(context context.Context, req *pb.RequestCodReserva) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-externo-by-codReserva")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioExterno, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoByCodReserva(context, req.GetCodReserva())
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao buscar o usuário externo pelo código de reserva "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscado o usuário externo pelo código de reserva",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioExterno", usuarioExterno),
	)

	return usuarioExterno, nil
}

// Função para criar um novo usuário externo
func (usuarioExternoServer *UsuariosExternosServer) CreateUsuarioExterno(context context.Context, req *pb.UsuarioPerfis) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-usuario-externo")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioExternoCriado, erroService := usuarioExternoServer.usuarioExternoService.CreateUsuarioExterno(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Erro ao criar o usuário externo "+erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo usuário externo",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioExterno", usuarioExternoCriado),
	)

	return usuarioExternoCriado, nil
}

// Função para criar um novo usuário externo
func (usuarioExternoServer *UsuariosExternosServer) CreateUsuarioExternoTeste(context context.Context, req *pb.UsuarioPerfis) (*pb.UsuarioPerfis, error) {
	usuarioExternoCriado, erroService := usuarioExternoServer.usuarioExternoService.CreateUsuarioExterno(context, req)
	if erroService.Erro != nil {
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	return usuarioExternoCriado, nil
}

// Função para atualizar um usuário externo já existente no banco
func (usuarioExternoServer *UsuariosExternosServer) UpdateUsuarioExterno(context context.Context, usuarioExterno *pb.UsuarioPerfis) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-usuario-externo")
	if retornoMiddleware.Erro != nil {
		return nil, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioExternoAntigo, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoById(context, usuarioExterno.GetUsuarioExterno().GetId())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	usuarioExternoAtualizado, erroService := usuarioExternoServer.usuarioExternoService.UpdateUsuarioExterno(context, usuarioExterno, usuarioExternoAntigo)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizado um usuário externo existente",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioExternoAntigo", usuarioExternoAntigo),
		zap.Any("usuarioExternoAtualizado", usuarioExternoAtualizado),
	)

	return usuarioExternoAtualizado, nil
}

// Função para o usuário alterar a própria senha
func (usuarioExternoServer *UsuariosExternosServer) AlterarSenhaUsuario(context context.Context, req *pb.RequestAlterarSenhaUsuario) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-alterar-propria-senha")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id não enviado")
	}

	usuarioASerAlterado, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	login := pb.LoginUsuario{
		Email: usuarioASerAlterado.UsuarioExterno.Email,
		Senha: req.SenhaAntiga,
	}

	usuarioLogado, erroService := usuarioExternoServer.usuarioExternoService.Login(context, &login)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = usuarioExternoServer.usuarioExternoService.AtualizarSenha(context, usuarioLogado.UsuarioExterno.Email, req.SenhaNova)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Alterada a senha de um usuário existente por meios próprios",
		zap.Any("usuario", usuarioSolicitante.UsuarioExterno),
		zap.Any("usuarioModificado", usuarioASerAlterado),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para desativar um usuário externo existente no banco
func (usuarioExternoServer *UsuariosExternosServer) DesativarUsuarioExterno(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-desativar-usuario-externo")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	usuarioExterno, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	desativado, erroService := usuarioExternoServer.usuarioExternoService.DesativarUsuarioExterno(context, id)
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

// Função para restaurar um usuário externo existente no banco
func (usuarioExternoServer *UsuariosExternosServer) RestaurarUsuarioExterno(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioExternoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-restaurar-usuario-externo")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetId()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "Id enviado não é válido ou não foi enviado")
	}

	usuarioExterno, erroService := usuarioExternoServer.usuarioExternoService.FindUsuarioExternoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.UsuarioExterno))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	restaurado, erroService := usuarioExternoServer.usuarioExternoService.RestaurarUsuarioExterno(context, id)
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

// Função para realizar o login na aplicação
func (usuarioExternoServer *UsuariosExternosServer) Login(context context.Context, req *pb.LoginUsuario) (*pb.UsuarioPerfis, error) {
	usuarioLogado, erroService := usuarioExternoServer.usuarioExternoService.Login(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Tentativa de login na API", zap.NamedError("err", erroService.Erro), zap.Any("email", req.GetEmail()))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Feito login na API",
		zap.Any("usuario", usuarioLogado.UsuarioExterno),
	)

	return usuarioLogado, nil
}

// Função para envio do token de reset de senha
func (usuarioExternoServer *UsuariosExternosServer) TokenResetSenha(context context.Context, req *pb.EmailReset) (*pb.ResponseTokenResetSenha, error) {
	token, err := usuarioExternoServer.createTokenReset(req)
	if err != nil {
		logger.Logger.Error(err.Error(), zap.NamedError("err", err), zap.Any("email", req.GetEmail()))
		return &pb.ResponseTokenResetSenha{}, status.Errorf(codes.Internal, err.Error())
	}

	erroService := usuarioExternoServer.usuarioExternoService.TokenResetSenha(context, token, req.GetEmail())
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
func (usuarioExternoServer *UsuariosExternosServer) ResetSenha(context context.Context, req *pb.ResetSenhaUsuario) (*pb.ResponseBool, error) {
	email, err := usuarioExternoServer.validarToken(req.Token)
	if err != nil {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.Unauthenticated, "Token inválido")
	}

	erroService := usuarioExternoServer.usuarioExternoService.AtualizarSenha(context, email, req.GetSenhaNova())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("email", email))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Resetada e alterada a senha de usuário",
		zap.Any("email", email),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// createToken cria um token JWT para um usuário autenticado para fins de reset da senha do mesmo.
func (usuarioExternoServer *UsuariosExternosServer) createTokenReset(emailReset *pb.EmailReset) (string, error) {
	claims := CustomClaimsResetSenha{
		Email:          emailReset.Email,
		UsuarioInterno: false,
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

	tokenString, err := token.SignedString(usuarioExternoServer.chaveSecreta) // Assina o token com a chave secreta
	if err != nil {
		logger.Logger.Error("Falha na criação do token de reset de senha", zap.NamedError("err", err), zap.Any("email", emailReset.GetEmail()))
		return "", err // Retorna erro se falhar
	}

	logger.Logger.Info("Criado um novo token de reset de senha",
		zap.Any("email", emailReset.GetEmail()),
	)

	return tokenString, nil // Retorna o token JWT gerado
}

// Função para verificar a validade de um token JWT
func (usuarioExternoServer *UsuariosExternosServer) validarToken(tokenString string) (string, error) {
	// Analisa o token JWT usando a chave secreta
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsResetSenha{}, func(token *jwt.Token) (interface{}, error) {
		return usuarioExternoServer.chaveSecreta, nil
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
