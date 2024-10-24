package middlewares

import (
	"context"
	"errors"
	"os"
	"si-admin/app/configuration/logger"
	"si-admin/app/model/erros"
	"si-admin/app/model/grpc"
	"si-admin/app/service"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var secretKey []byte // Variável global para armazenar a chave secreta usada na assinatura dos tokens JWT

// Struct para capturar as informações do token JWT
type CustomClaims struct {
	IdUsuario            int32 // Id do usuário proprietário do token
	UsuarioInterno       bool
	jwt.RegisteredClaims // Struct que contém os claims padrão do JWT
}

// Estrutura do middleware para permissões
type PermissoesMiddleware struct {
	usuarioInternoService *service.UsuarioInternoService // Serviço para manipulação de usuários
	usuarioExternoService *service.UsuarioExternoService // Serviço para manipulação de usuários
	perfilService         *service.PerfilService         // Serviço para manipulação de perfis
}

// Função para criar uma nova instância do middleware de permissões
func NewPermissoesMiddleware(usuarioInternoService *service.UsuarioInternoService, perfilService *service.PerfilService, usuarioExternoService *service.UsuarioExternoService) *PermissoesMiddleware {
	return &PermissoesMiddleware{
		usuarioInternoService: usuarioInternoService,
		usuarioExternoService: usuarioExternoService,
		perfilService:         perfilService,
	}
}

// Função para verificar as permissões de acesso a uma rota específica
func (permissoesMiddleware *PermissoesMiddleware) PermissaoMiddleware(context context.Context, rota string) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	var err error
	// Lê a chave secreta do arquivo
	secretKey, err = os.ReadFile("./jwt/jwt_secret_key.txt")
	if err != nil {
		panic("Erro ao ler a chave secreta") // Interrompe a execução se houver erro ao ler a chave
	}

	// Obtém o token de autorização dos metadados da requisição
	metadata, okMetadata := metadata.FromIncomingContext(context)
	if !okMetadata {
		erro := errors.New("Não foi enviado nenhum metadata")
		logger.Logger.Error("Tentativa de acesso sem nenhum metadata", zap.NamedError("err", erro), zap.String("rota", rota))
		return nil, erros.ErroStatus{
			Status: codes.Unauthenticated,
			Erro:   erro,
		}
	}

	itens := metadata["authorization"]
	if len(itens) == 0 {
		erro := errors.New("Não foi enviado o token de autorização")
		logger.Logger.Error("Tentativa de acesso sem token de autorização", zap.NamedError("err", erro), zap.String("rota", rota))
		return nil, erros.ErroStatus{
			Status: codes.Unauthenticated,
			Erro:   erro,
		}
	}

	// Remove o prefixo "Bearer " do token
	tokenString := strings.TrimPrefix(itens[0], "Bearer ")
	if tokenString == "" {
		erro := errors.New("Não foi enviado o token de autorização")
		logger.Logger.Error("Tentativa de acesso sem token de autorização", zap.NamedError("err", erro), zap.String("rota", rota))
		return nil, erros.ErroStatus{
			Status: codes.Unauthenticated,
			Erro:   erro,
		}
	}

	props, errVerificacao := verifyToken(tokenString) // Verifica a validade do token e retorna o id de usuário contido nele
	if errVerificacao != nil {
		erro := errors.New("Token inválido")
		logger.Logger.Error("Tentativa de acesso com um token inválido", zap.NamedError("err", erro), zap.String("rota", rota))
		return nil, erros.ErroStatus{
			Status: codes.Unauthenticated,
			Erro:   erro,
		}
	}

	var perfisVinculados []*grpc.Perfil
	var usuarioPerfil *grpc.UsuarioPerfis

	if props.UsuarioInterno {
		// Busca pelo usuário proprietário do token baseado no id que constar no token
		usuario, erroService := permissoesMiddleware.usuarioInternoService.FindUsuarioInternoById(context, props.IdUsuario)
		if erroService.Erro != nil {
			return nil, erros.ErroStatus{
				Status: erroService.Status,
				Erro:   erroService.Erro,
			}
		}

		logger.Logger.Info("Solicitado acesso à API", zap.Any("usuario", usuario.UsuarioInterno), zap.String("rota", rota))

		// Verifica a flag SenhaAtualizada do usuário
		// Caso essa flag seja true, significa que o usuário trocou a senha dele, invalidando esse token e interrompendo a execução da request
		// Essa flag é colocada como true sempre que a senha é atualizada, seja por reset, modificação própria ou do admin
		// Sempre que um token é criado para aquele usuário, essa flag é definida como false
		// Ou seja, quando a senha for modificada, esse middleware vai identificar e impedir o acesso
		// Obrigando o usuário a gerar um novo token com a nova senha
		if usuario.GetUsuarioInterno().GetSenhaAtualizada() {
			erro := errors.New("Token inválido")
			logger.Logger.Error("Tentativa de acesso com um token inválido", zap.NamedError("err", erro), zap.String("rota", rota))
			return nil, erros.ErroStatus{
				Status: codes.Unauthenticated,
				Erro:   erro,
			}
		}

		if !usuario.GetUsuarioInterno().GetAtivo() {
			erro := errors.New("Usuário desativado")
			logger.Logger.Error("Tentativa de acesso com um usuário desativado", zap.NamedError("err", erro), zap.String("rota", rota))
			return nil, erros.ErroStatus{
				Status: codes.Unauthenticated,
				Erro:   erro,
			}
		}

		perfisVinculados = usuario.GetPerfis()
		usuarioPerfil = usuario
	} else {
		// Busca pelo usuário proprietário do token baseado no id que constar no token
		usuario, erroService := permissoesMiddleware.usuarioExternoService.FindUsuarioExternoById(context, props.IdUsuario)
		if erroService.Erro != nil {
			return nil, erros.ErroStatus{
				Status: erroService.Status,
				Erro:   erroService.Erro,
			}
		}

		logger.Logger.Info("Solicitado acesso à API", zap.Any("usuario", usuario.UsuarioExterno), zap.String("rota", rota))

		// Verifica a flag SenhaAtualizada do usuário
		// Caso essa flag seja true, significa que o usuário trocou a senha dele, invalidando esse token e interrompendo a execução da request
		// Essa flag é colocada como true sempre que a senha é atualizada, seja por reset, modificação própria ou do admin
		// Sempre que um token é criado para aquele usuário, essa flag é definida como false
		// Ou seja, quando a senha for modificada, esse middleware vai identificar e impedir o acesso
		// Obrigando o usuário a gerar um novo token com a nova senha
		if usuario.GetUsuarioExterno().GetSenhaAtualizada() {
			erro := errors.New("Token inválido")
			logger.Logger.Error("Tentativa de acesso com um token inválido", zap.NamedError("err", erro), zap.String("rota", rota))
			return nil, erros.ErroStatus{
				Status: codes.Unauthenticated,
				Erro:   erro,
			}
		}

		if !usuario.GetUsuarioExterno().GetAtivo() {
			erro := errors.New("Usuário desativado")
			logger.Logger.Error("Tentativa de acesso com um usuário desativado", zap.NamedError("err", erro), zap.String("rota", rota))
			return nil, erros.ErroStatus{
				Status: codes.Unauthenticated,
				Erro:   erro,
			}
		}

		perfisVinculados = usuario.GetPerfis()
		usuarioPerfil = usuario
	}

	// Verifica se algum dos perfis do usuário tem a permissão necessária para acessar a rota
	for i := range perfisVinculados {
		permissoesVinculadas, erroService := permissoesMiddleware.perfilService.GetPermissoesVinculadas(context, perfisVinculados[i].GetId())
		if erroService.Erro != nil {
			return nil, erros.ErroStatus{
				Status: erroService.Status,
				Erro:   erroService.Erro,
			}
		}

		// Verifica se a permissão necessária está presente nas permissões do perfil
		for _, permissao := range permissoesVinculadas {
			if rota == permissao.Nome {
				logger.Logger.Info("Acessada a rota", zap.Any("usuarioExterno", usuarioPerfil.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioPerfil.GetUsuarioInterno()), zap.String("rota", rota))
				return usuarioPerfil, erros.ErroStatus{}
			}
		}
	}

	// Retorna um erro 401 (Unauthorized) se o usuário não tiver a permissão necessária
	erro := errors.New("Permissões insuficientes")
	logger.Logger.Error("Tentativa de acesso a uma rota sem ter a permissão para tal", zap.Any("usuarioExterno", usuarioPerfil.GetUsuarioExterno()), zap.Any("usuarioInterno", usuarioPerfil.GetUsuarioInterno()), zap.String("rota", rota))

	return nil, erros.ErroStatus{
		Status: codes.PermissionDenied,
		Erro:   erro,
	}
}

// Função para verificar a validade de um token JWT
func verifyToken(tokenString string) (*CustomClaims, error) {
	// Analisa o token JWT usando a chave secreta
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err // Retorna erro se o token não puder ser analisado
	}

	if !token.Valid {
		return nil, errors.New("Token inválido") // Retorna erro se o token não for válido
	}

	// Extrai as claims personalizadas do token
	props, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("Propriedades inválidas do token") // Retorna erro se as claims não puderem ser extraídas
	}

	return props, nil // Retorna o id do usuário
}
