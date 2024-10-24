package service

import (
	"context"
	"encoding/base64"
	"errors"
	"si-admin/app/helpers"
	"si-admin/app/model/erros"
	pb "si-admin/app/model/grpc"
	"si-admin/app/model/repositories"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
)

// Estrutura que define o serviço de usuário, contendo os repositórios necessários
type UsuarioInternoService struct {
	perfilRepository         repositories.PerfilRepository
	usuarioPerfilRepository  repositories.UsuarioPerfilRepository
	usuarioInternoRepository repositories.UsuarioInternoRepository
}

// Função para criar uma nova instância de UsuarioInternoService com os repositórios fornecidos
func NewUsuarioInternoService(perfilRepository repositories.PerfilRepository, usuarioPerfilRepository repositories.UsuarioPerfilRepository, usuarioInternoRepository repositories.UsuarioInternoRepository) *UsuarioInternoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &UsuarioInternoService{
		perfilRepository:         perfilRepository,
		usuarioPerfilRepository:  usuarioPerfilRepository,
		usuarioInternoRepository: usuarioInternoRepository,
	}
}

// Função para buscar todos os usuários
// Como um usuário pode estar vinculado a vários perfis,
// Nessa função, os usuários são retornados sem as perfis vinculados
func (usuarioService *UsuarioInternoService) FindAllUsuariosInternos(context context.Context, req *pb.RequestAllPaginado) (*pb.ListaUsuariosInternos, erros.ErroStatus) {
	if req.GetTamanhoPagina() == 0 {
		req.TamanhoPagina = 10
	}

	tUsuariosInternos, err := usuarioService.usuarioInternoRepository.FindAll(context, req.GetCursor(), (req.GetTamanhoPagina() + 1))
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma cidade, retorna code NotFound
	if len(tUsuariosInternos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma cidade encontrada"),
		}
	}

	var proximoCursor int32
	temMais := false

	if len(tUsuariosInternos) == (int(req.GetTamanhoPagina()) + 1) {
		temMais = true
	}

	var pbCidades []*pb.UsuarioInterno
	for i, usuario := range tUsuariosInternos {
		if i == (int(req.GetTamanhoPagina())) {
			break
		}

		pbCidades = append(pbCidades, helpers.TUsuarioInternoToPb(usuario))
		proximoCursor = usuario.ID
	}

	return &pb.ListaUsuariosInternos{UsuariosInternos: pbCidades, Meta: &pb.Meta{ProximoCursor: proximoCursor, TamanhoPagina: req.GetTamanhoPagina(), TemMais: temMais}}, erros.ErroStatus{}
}

// Função para buscar um usuário por ID
// Diferente da busca por todos, aqui a aplicação retorna o usuário com todos os perfis vinculados
func (usuarioService *UsuarioInternoService) FindUsuarioInternoById(context context.Context, id int32) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuario, err := usuarioService.usuarioInternoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	perfis, erroPerfis := usuarioService.GetPerfisVinculados(context, usuario.ID)
	if erroPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	retorno := pb.UsuarioPerfis{
		UsuarioInterno: helpers.TUsuarioInternoToPb(usuario),
		Perfis:         pbPerfis,
	}

	return &retorno, erros.ErroStatus{}
}

// Função para buscar perfis vinculados a um usuário, filtrando apenas os perfis ativos
func (usuarioService *UsuarioInternoService) GetPerfisVinculados(context context.Context, id int32) ([]pb.Perfil, erros.ErroStatus) {
	usuarioPerfis, errBuscaPerfis := usuarioService.usuarioPerfilRepository.FindByUsuarioInterno(context, id)
	if errBuscaPerfis != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if errBuscaPerfis.Error() == "no rows in result set" {
			return []pb.Perfil{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaPerfis,
		}
	}

	// Filtra apenas os perfis que estão vinculados e ativos
	var perfis []pb.Perfil
	for _, usuarioPerfil := range usuarioPerfis {
		perfilEncontrado, err := usuarioService.perfilRepository.FindByID(context, usuarioPerfil.PerfilID)
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		if perfilEncontrado.Ativo.Bool {
			perfis = append(perfis, *helpers.TPerfToPb(perfilEncontrado))
		}
	}

	return perfis, erros.ErroStatus{}
}

// Função de login, que verifica o email, a senha do usuário e se ele está ativo ou não
// Retorna o usuário com seus perfis vinculados (UsuarioPerfis)
func (usuarioService *UsuarioInternoService) Login(context context.Context, loginUsuarioInterno *pb.LoginUsuario) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuarioBanco, err := usuarioService.usuarioInternoRepository.FindByEmail(context, loginUsuarioInterno.Email)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code Unauthenticated
		// Dessa forma, o usuário que está tentando acessar não consegue saber se o que ele errou foi o e-mail ou a senha
		// Assim acredito que fique mais seguro contra tentativas maliciosas de acesso
		if err.Error() == "no rows in result set" {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Unauthenticated,
				Erro:   errors.New("E-mail e/ou senha incorretos"),
			}
		}

		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	if !usuarioBanco.Ativo.Bool {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.InvalidArgument,
			Erro:   errors.New("Usuário está desativado"),
		}
	}

	perfisVinculados, erroGetPerfis := usuarioService.GetPerfisVinculados(context, usuarioBanco.ID)
	if erroGetPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroGetPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfisVinculados {
		pbPerfis = append(pbPerfis, &perfisVinculados[i])
	}

	usuarioPerfil := new(pb.UsuarioPerfis)
	usuarioPerfil.UsuarioInterno = helpers.TUsuarioInternoToPb(usuarioBanco)
	usuarioPerfil.Perfis = pbPerfis

	// Validação da senha codificada do banco
	senhaHash, err := base64.StdEncoding.DecodeString(usuarioBanco.Senha)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	erroValidacao := bcrypt.CompareHashAndPassword(senhaHash, []byte(loginUsuarioInterno.Senha))
	if erroValidacao != nil {
		// Caso a senha esteja incorreta, não informa que é a senha o problema, mas sim um erro genérico de login
		// Dessa forma, o usuário que está tentando acessar não consegue saber se o que ele errou foi o e-mail ou a senha
		// Assim acredito que fique mais seguro contra tentativas maliciosas de acesso
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Unauthenticated,
			Erro:   errors.New("E-mail e/ou senha incorretos"),
		}
	} else {
		return usuarioPerfil, erros.ErroStatus{}
	}
}

// Essa função insere na base de dados o token criado para o reset da senha do usuário em caso de esquecimento
// Retorna apenas um erro caso aconteça algum
func (usuarioService *UsuarioInternoService) TokenResetSenha(context context.Context, token string, email string) erros.ErroStatus {
	usuarioBanco, err := usuarioService.usuarioInternoRepository.FindByEmail(context, email)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Salva o usuário atualizado no repositório
	_, err = usuarioService.usuarioInternoRepository.SetTokenResetSenha(context, token, usuarioBanco.ID)
	if err != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return erros.ErroStatus{}
}

// Função para atualizar a senha de um usuário
// Retorna apenas um erro caso aconteça algum
func (usuarioService *UsuarioInternoService) AtualizarSenha(context context.Context, email string, senhaNova string) erros.ErroStatus {
	usuarioEncontrado, err := usuarioService.usuarioInternoRepository.FindByEmail(context, email)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Criptografia da senha
	senhaEmBytes := []byte(senhaNova)
	senhaCriptografada, erroCriptografia := bcrypt.GenerateFromPassword(senhaEmBytes, 10)
	if erroCriptografia != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroCriptografia,
		}
	}

	// Salva o usuário atualizado no repositório
	_, err = usuarioService.usuarioInternoRepository.UpdateSenha(context, base64.StdEncoding.EncodeToString(senhaCriptografada), usuarioEncontrado.ID)
	if err != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return erros.ErroStatus{}
}

// Função para clonar um usuário, mantendo as permissões, mas retornando um novo usuário vazio
func (usuarioService *UsuarioInternoService) CloneUsuarioInterno(context context.Context, id int32) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuarioPerfil, err := usuarioService.FindUsuarioInternoById(context, id)
	if err.Erro != nil {
		return &pb.UsuarioPerfis{}, err
	}

	// Define o usuário clonado como vazio
	usuarioPerfil.UsuarioInterno = &pb.UsuarioInterno{}

	return usuarioPerfil, erros.ErroStatus{}
}

// Função para criar um novo usuário, verificando se o e-mail já existe e criptografando a senha
func (usuarioService *UsuarioInternoService) CreateUsuarioInterno(context context.Context, usuarioPerfilRecebido *pb.UsuarioPerfis) (*pb.UsuarioPerfis, erros.ErroStatus) {
	// Busca um usuário pelo e-mail enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	_, err := usuarioService.usuarioInternoRepository.FindByEmail(context, usuarioPerfilRecebido.GetUsuarioInterno().GetEmail())
	if err == nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("E-mail já está sendo utilizado"),
		}
	}

	// Criptografia da senha
	senhaEmBytes := []byte(usuarioPerfilRecebido.GetUsuarioInterno().GetSenha())
	senhaCriptografada, erroCriptografia := bcrypt.GenerateFromPassword(senhaEmBytes, 10)
	if erroCriptografia != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroCriptografia,
		}
	}

	// Cria o objeto CreateUsuarioInternoParams gerado pelo sqlc para gravação no banco de dados
	usuarioCreate := db.CreateUsuarioInternoParams{
		Nome:            usuarioPerfilRecebido.GetUsuarioInterno().GetNome(),
		Email:           usuarioPerfilRecebido.GetUsuarioInterno().GetEmail(),
		Senha:           base64.StdEncoding.EncodeToString(senhaCriptografada),
		Ativo:           pgtype.Bool{Bool: true, Valid: true},
		TokenResetSenha: pgtype.Text{String: usuarioPerfilRecebido.GetUsuarioInterno().GetTokenResetSenha(), Valid: true},
		SenhaAtualizada: pgtype.Bool{Bool: true, Valid: true},
	}

	// Cria o usuário no repositório
	usuario, err := usuarioService.usuarioInternoRepository.Create(context, usuarioCreate)
	usuarioPerfilRecebido.UsuarioInterno = helpers.TUsuarioInternoToPb(usuario)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Salva a relação entre o usuário e cada perfil
	for i, perfil := range usuarioPerfilRecebido.GetPerfis() {
		perfil, err := usuarioService.perfilRepository.FindByID(context, perfil.Id)
		if err != nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo do perfil na lista de perfis do objeto recebido
		// Como esse é o objeto retornado, os perfis retornam completos e prontos para serem usados
		usuarioPerfilRecebido.Perfis[i] = helpers.TPerfToPb(perfil)

		// Cria o objeto CreateUsuarioPerfilParams gerado pelo sqlc para gravação no banco de dados
		usuarioPerfil := db.CreateUsuarioPerfilParams{
			UsuarioInternoID: pgtype.Int4{Int32: usuario.ID, Valid: true},
			PerfilID:         perfil.ID,
			DataHora:         pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		usuarioService.usuarioPerfilRepository.Create(context, usuarioPerfil)
	}

	return usuarioPerfilRecebido, erros.ErroStatus{}
}

// Função para atualizar um usuário, verificando o email e atualizando perfis relacionados
func (usuarioService *UsuarioInternoService) UpdateUsuarioInterno(context context.Context, usuarioPerfil *pb.UsuarioPerfis, usuarioAntigo *pb.UsuarioPerfis) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuarioRecebido := usuarioPerfil.UsuarioInterno
	usuarioBanco := usuarioAntigo.GetUsuarioInterno()
	// Verifica se o e-mail foi modificado e, se sim, verifica se já existe outro registro com o mesmo e-mail
	// Em caso positivo, retorna code AlreadyExists
	if usuarioBanco.Email != usuarioRecebido.Email {
		_, err := usuarioService.usuarioInternoRepository.FindByEmail(context, usuarioRecebido.Email)
		if err == nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("E-mail já está sendo utilizado"),
			}
		}
	}

	// Cria o objeto UpdateUsuarioInternoParams gerado pelo sqlc para gravação no banco de dados
	usuarioUpdate := db.UpdateUsuarioInternoParams{
		ID:    usuarioBanco.Id,
		Nome:  usuarioRecebido.Nome,
		Email: usuarioRecebido.Email,
	}

	// Salva o usuário atualizado no repositório
	usuarioAtualizado, err := usuarioService.usuarioInternoRepository.Update(context, usuarioUpdate)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Deleta todas as relações antigas entre o usuário e seus perfis
	registrosExistentes, errBuscaPerfis := usuarioService.usuarioPerfilRepository.FindByUsuarioInterno(context, usuarioAtualizado.ID)
	if errBuscaPerfis != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaPerfis,
		}
	}

	for i := 0; i < len(registrosExistentes); i++ {
		err := usuarioService.usuarioPerfilRepository.Delete(context, registrosExistentes[i].ID)
		if err != nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	}

	// Cria novas relações entre o perfil atualizado e as novas permissões
	perfis := usuarioPerfil.Perfis

	for i, perfil := range perfis {
		tPerfil, err := usuarioService.perfilRepository.FindByID(context, perfil.Id)
		if err != nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo do perfil na lista de perfis do objeto recebido
		// Como esse é o objeto retornado, os perfis retornam completos e prontos para serem usados
		perfis[i] = helpers.TPerfToPb(tPerfil)

		// Cria o objeto CreateUsuarioPerfilParams gerado pelo sqlc para gravação no banco de dados
		usuarioPerfil := db.CreateUsuarioPerfilParams{
			UsuarioInternoID: pgtype.Int4{Int32: usuarioAtualizado.ID, Valid: true},
			PerfilID:         perfil.Id,
			DataHora:         pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		usuarioService.usuarioPerfilRepository.Create(context, usuarioPerfil)
	}

	usuarioPerfilRetorno := pb.UsuarioPerfis{
		UsuarioInterno: helpers.TUsuarioInternoToPb(usuarioAtualizado),
		Perfis:         perfis,
	}

	return &usuarioPerfilRetorno, erros.ErroStatus{}
}

// / Função para desativar um usuário externo pelo Id
func (usuarioService *UsuarioInternoService) DesativarUsuarioInterno(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Desativa o usuário externo no repositório pelo Id
	desativados, err := usuarioService.usuarioInternoRepository.Desativar(context, id)

	// Caso ocorra erro na desativação
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo desativados indica o número de linhas desativadas. Se for 0, nenhum usuário externo foi desativado, pois não existia
	if desativados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum usuário externo encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}

// Função para restaurar um usuário externo pelo Id
func (usuarioService *UsuarioInternoService) RestaurarUsuarioInterno(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Desativa o usuário externo no repositório pelo Id
	restaurados, err := usuarioService.usuarioInternoRepository.Restaurar(context, id)

	// Caso ocorra erro na restauração
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo restaurados indica o número de linhas restauradas. Se for 0, nenhum usuário externo foi restaurado, pois não existia
	if restaurados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum usuário externo encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}

// Essa função serve apenas para definir a flag `SenhaAtualizada` como false
// Dessa forma o usuário consegue passar pelo middleware e acessar a API
// Essa função é chamada sempre que um usuário faz login
func (usuarioService *UsuarioInternoService) SenhaAtualizada(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Salva o usuário atualizado no repositório
	atualizados, err := usuarioService.usuarioInternoRepository.SetSenhaAtualizada(context, id)

	// Caso ocorra erro na restauração
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo atualizados indica o número de linhas atualizados. Se for 0, nenhum usuário interno foi atualizados, pois não existia
	if atualizados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum usuário interno encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}
