package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"si-admin/app/controller"
	"si-admin/app/controller/middlewares"
	pb "si-admin/app/model/grpc"
	"si-admin/app/model/repositories"
	"si-admin/app/model/repositories/sqlc/repositoryIMPL"
	"si-admin/app/service"
	"si-admin/db"

	"google.golang.org/grpc"
)

type ParametroKey string

func main() {
	ctx := context.Background()
	// Configuração de chave secreta do token JWT
	if os.Getenv("GENERATE_KEY") == "true" {
		if err := GeraChaveSecreta(); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gerar a chave secreta: %v\n", err)
			os.Exit(1)
		}
	}

	chaveSecreta, err := os.ReadFile("./jwt/jwt_secret_key.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a chave secreta: %v\n", err)
		os.Exit(1)
	}

	db := db.CreateConnection()
	defer db.Close()

	sqlcQueries := repositoryIMPL.New(db)

	// Repositórios
	// Aqui são feitas as implementações das funções que realmente interagirão com o banco de dados
	usuarioExternoRepository := repositories.NewUsuarioExternoRepository(sqlcQueries)
	perfilPermissaoRepository := repositories.NewPerfilPermissaoRepository(sqlcQueries)
	perfilRepository := repositories.NewPerfilRepository(sqlcQueries)
	permissaoRepository := repositories.NewPermissaoRepository(sqlcQueries)
	usuarioPerfilRepository := repositories.NewUsuarioPerfilRepository(sqlcQueries)
	usuarioInternoRepository := repositories.NewUsuarioInternoRepository(sqlcQueries)

	// Serviços
	// Aqui é a camada da minha regra de negócio, todas as validações, modificações e adaptações dos dados são feitas aqui
	usuarioExternoService := service.NewUsuarioExternoService(perfilRepository, usuarioPerfilRepository, usuarioExternoRepository)
	perfilService := service.NewPerfilService(perfilPermissaoRepository, permissaoRepository, usuarioPerfilRepository, perfilRepository, usuarioInternoRepository, usuarioExternoRepository)
	permissaoService := service.NewPermissaoService(perfilPermissaoRepository, permissaoRepository)
	usuarioInternoService := service.NewUsuarioInternoService(perfilRepository, usuarioPerfilRepository, usuarioInternoRepository)

	// Middleware de acesso
	permissaoMiddleware := middlewares.NewPermissoesMiddleware(usuarioInternoService, perfilService, usuarioExternoService)
	// O "middleware" acima está sendo utilizado atualmente como uma chamada no início da execução de cada uma das funções dos servers/controllers
	// A função chama a função passando o contexto e a permissão específica daquele comando
	// Dessa forma, caso o usuário tenha as permissões necessárias, é retornado um erro nulo e continada a requisição
	// Caso o usuário não tenha as permissões necessárias, é abortada a requisição passando o erro que impediu esse acesso

	// Servers gRPC
	// Aqui é a camada de controle, aqui não temos regras de negócio, apenas a comunicação efetiva com o cliente
	usuarioExternoServer := controller.NewUsuarioExternoServer(usuarioExternoService, permissaoMiddleware, chaveSecreta)
	perfisServer := controller.NewPerfisServer(perfilService, permissaoMiddleware)
	permissoesServer := controller.NewPermissoesServer(permissaoService, permissaoMiddleware)
	usuarioInternoServer := controller.NewUsuariosInternosServer(usuarioInternoService, permissaoMiddleware, chaveSecreta)

	// Criado o listener para a porta da aplicação
	lis, err := net.Listen("tcp", ":8601")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
		return
	}

	// Carrega o certificado e a chave para TLS
	// creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	// if err != nil {
	// 	log.Fatalf("Falha ao carregar credenciais de TLS: %v", err)
	// 	return
	// }

	// Configurações do servidor gRPC com TLS
	serverGrpc := grpc.NewServer()

	// Registro dos servidores/controllers
	pb.RegisterUsuariosExternosServer(serverGrpc, usuarioExternoServer)
	pb.RegisterPerfisServer(serverGrpc, perfisServer)
	pb.RegisterPermissoesServer(serverGrpc, permissoesServer)
	pb.RegisterUsuariosInternosServer(serverGrpc, usuarioInternoServer)

	// Feitas as configurações e inicializações, vamos dar início ao processo de inicialização da aplicação

	// Criação do primeiro usuário (admin) baseado no json abaixo
	_, erroService := usuarioInternoService.FindUsuarioInternoById(ctx, 1)
	if erroService.Erro != nil {
		jsonFile, err := os.Open("./seedUsuarioAdmin.json")
		if err != nil {
			log.Fatalf("Falha ao criar o usuário admin: %v", err.Error())
			os.Exit(1)
		}

		defer jsonFile.Close()
		byteValueJson, _ := io.ReadAll(jsonFile)
		objUsuario := pb.UsuarioPerfis{}
		json.Unmarshal(byteValueJson, &objUsuario)
		_, erroCriacaoUsuario := usuarioInternoService.CreateUsuarioInterno(ctx, &objUsuario)
		if erroCriacaoUsuario.Erro != nil {
			if erroService.Erro != nil { // Se houver erro na criação, retorna erro
				log.Fatalf("Falha ao criar o usuário admin: %v", err.Error())
				os.Exit(1)
			}
		}
	}

	log.Println("Iniciando servidor gRPC na porta 8601...")

	if err := serverGrpc.Serve(lis); err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
}
