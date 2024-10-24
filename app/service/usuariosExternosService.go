package service

import (
	"context"
	"encoding/base64"
	"errors"
	"si-admin/app/helpers"
	"si-admin/app/model/erros"
	"si-admin/app/model/grpc"
	pb "si-admin/app/model/grpc"
	"si-admin/app/model/repositories"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
)

// Estrutura de serviço para gerenciar operações relacionadas às clientes
type UsuarioExternoService struct {
	perfilRepository         repositories.PerfilRepository
	usuarioPerfilRepository  repositories.UsuarioPerfilRepository
	usuarioExternoRepository repositories.UsuarioExternoRepository
}

// Função para criar uma nova instância de UsuarioExternoService com o repositório necessário
func NewUsuarioExternoService(perfilRepository repositories.PerfilRepository, usuarioPerfilRepository repositories.UsuarioPerfilRepository, usuarioRepository repositories.UsuarioExternoRepository) *UsuarioExternoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &UsuarioExternoService{
		perfilRepository:         perfilRepository,
		usuarioPerfilRepository:  usuarioPerfilRepository,
		usuarioExternoRepository: usuarioRepository,
	}
}

// Função para buscar um usuário externo pelo Id
func (usuarioExternoService *UsuarioExternoService) FindUsuarioExternoById(context context.Context, id int32) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	// Busca o usuario no repositório pelo Id
	usuarioExterno, err := usuarioExternoService.usuarioExternoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum usuário externo, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return nil, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário externo não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	perfis, erroPerfis := usuarioExternoService.GetPerfisVinculados(context, usuarioExterno.ID)
	if erroPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	retorno := pb.UsuarioPerfis{
		UsuarioExterno: helpers.TUsuarioExternoToPb(usuarioExterno),
		Perfis:         pbPerfis,
	}

	return &retorno, erros.ErroStatus{}
}

// Função para buscar um usuário externo pelo Id externo
func (usuarioExternoService *UsuarioExternoService) FindUsuarioExternoByIdExterno(context context.Context, id int32) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	// Busca o usuario no repositório pelo Id
	usuarioExterno, err := usuarioExternoService.usuarioExternoRepository.FindByIDExterno(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum usuário externo, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return nil, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário externo não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	perfis, erroPerfis := usuarioExternoService.GetPerfisVinculados(context, usuarioExterno.ID)
	if erroPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	retorno := pb.UsuarioPerfis{
		UsuarioExterno: helpers.TUsuarioExternoToPb(usuarioExterno),
		Perfis:         pbPerfis,
	}

	return &retorno, erros.ErroStatus{}
}

// Função para buscar um usuário externo pelo nome
func (usuarioExternoService *UsuarioExternoService) FindUsuarioExternoByNome(context context.Context, req *grpc.RequestNome) (*grpc.ListaUsuariosExternos, erros.ErroStatus) {
	if req.GetTamanhoPagina() == 0 {
		req.TamanhoPagina = 10
	}

	// Busca o usuário externo no repositório pelo nome
	tUsuarioExternos, err := usuarioExternoService.usuarioExternoRepository.FindByNome(context, req.GetNome(), req.GetCursor(), (req.GetTamanhoPagina() + 1))
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum usuário externo, retorna code NotFound
	if len(tUsuarioExternos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum usuário externo encontrado"),
		}
	}

	var proximoCursor int32
	temMais := false

	if len(tUsuarioExternos) == (int(req.GetTamanhoPagina()) + 1) {
		temMais = true
	}

	var grpcUsuarioExternos []*grpc.UsuarioExterno
	for i, usuarioExterno := range tUsuarioExternos {
		if i == (int(req.GetTamanhoPagina())) {
			break
		}

		grpcUsuarioExternos = append(grpcUsuarioExternos, helpers.TUsuarioExternoToPb(usuarioExterno))
		proximoCursor = usuarioExterno.ID
	}

	return &grpc.ListaUsuariosExternos{UsuariosExternos: grpcUsuarioExternos, Meta: &grpc.Meta{ProximoCursor: proximoCursor, TamanhoPagina: req.GetTamanhoPagina(), TemMais: temMais}}, erros.ErroStatus{}
}

// Função para buscar um usuário externo pelo documento
func (usuarioExternoService *UsuarioExternoService) FindUsuarioExternoByDocumento(context context.Context, documento string) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	// Busca o usuario no repositório pelo nome
	usuarioExterno, err := usuarioExternoService.usuarioExternoRepository.FindByDocumento(context, documento)
	if err != nil {
		// Caso não seja encontrado nenhum usuário externo, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return nil, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário externo não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	perfis, erroPerfis := usuarioExternoService.GetPerfisVinculados(context, usuarioExterno.ID)
	if erroPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	retorno := pb.UsuarioPerfis{
		UsuarioExterno: helpers.TUsuarioExternoToPb(usuarioExterno),
		Perfis:         pbPerfis,
	}

	return &retorno, erros.ErroStatus{}
}

// Função para buscar um usuário externo pelo documento
func (usuarioExternoService *UsuarioExternoService) FindUsuarioExternoByEmail(context context.Context, email string) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	// Busca o usuário externo no repositório pelo nome
	usuarioExterno, err := usuarioExternoService.usuarioExternoRepository.FindByEmail(context, email)
	if err != nil {
		// Caso não seja encontrado nenhum usuário externo, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return nil, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário externo não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	perfis, erroPerfis := usuarioExternoService.GetPerfisVinculados(context, usuarioExterno.ID)
	if erroPerfis.Erro != nil {
		return nil, erroPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	retorno := pb.UsuarioPerfis{
		UsuarioExterno: helpers.TUsuarioExternoToPb(usuarioExterno),
		Perfis:         pbPerfis,
	}

	return &retorno, erros.ErroStatus{}
}

// Função para buscar um usuário externo pelo código de reserva
func (usuarioExternoService *UsuarioExternoService) FindUsuarioExternoByCodReserva(context context.Context, codReserva string) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	// Busca o usuário externo no repositório pelo código de reserva
	usuario, err := usuarioExternoService.usuarioExternoRepository.FindByCodReserva(context, codReserva)
	if err != nil {
		// Caso não seja encontrado nenhum usuário externo, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return nil, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário externo não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	perfis, erroPerfis := usuarioExternoService.GetPerfisVinculados(context, usuario.ID)
	if erroPerfis.Erro != nil {
		return nil, erroPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	retorno := pb.UsuarioPerfis{
		UsuarioExterno: helpers.TUsuarioExternoToPb(usuario),
		Perfis:         pbPerfis,
	}

	return &retorno, erros.ErroStatus{}
}

// Função para buscar todas os usuários externo
func (usuarioExternoService *UsuarioExternoService) FindAllUsuariosExternos(context context.Context, req *grpc.RequestAllPaginado) (*grpc.ListaUsuariosExternos, erros.ErroStatus) {
	if req.GetTamanhoPagina() == 0 {
		req.TamanhoPagina = 10
	}

	// Busca todas os usuários externos no repositório
	usuariosExternos, err := usuarioExternoService.usuarioExternoRepository.FindAll(context, req.GetCursor(), (req.GetTamanhoPagina() + 1))
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum usuário externo, retorna code NotFound
	if len(usuariosExternos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum usuário externo encontrado"),
		}
	}

	var proximoCursor int32
	temMais := false

	if len(usuariosExternos) == (int(req.GetTamanhoPagina()) + 1) {
		temMais = true
	}

	var grpcUsuarioExternos []*grpc.UsuarioExterno
	for i, usuarioExterno := range usuariosExternos {
		if i == (int(req.GetTamanhoPagina())) {
			break
		}

		grpcUsuarioExternos = append(grpcUsuarioExternos, helpers.TUsuarioExternoToPb(usuarioExterno))
		proximoCursor = usuarioExterno.ID
	}

	return &grpc.ListaUsuariosExternos{UsuariosExternos: grpcUsuarioExternos, Meta: &grpc.Meta{ProximoCursor: proximoCursor, TamanhoPagina: req.GetTamanhoPagina(), TemMais: temMais}}, erros.ErroStatus{}
}

// Função para buscar perfis vinculados a um usuário, filtrando apenas os perfis ativos
func (usuarioExternoService *UsuarioExternoService) GetPerfisVinculados(context context.Context, id int32) ([]pb.Perfil, erros.ErroStatus) {
	usuarioPerfis, errBuscaPerfis := usuarioExternoService.usuarioPerfilRepository.FindByUsuarioExterno(context, id)
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
		perfilEncontrado, err := usuarioExternoService.perfilRepository.FindByID(context, usuarioPerfil.PerfilID)
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

// Essa função insere na base de dados o token criado para o reset da senha do usuário em caso de esquecimento
// Retorna apenas um erro caso aconteça algum
func (usuarioExternoService *UsuarioExternoService) TokenResetSenha(context context.Context, token string, email string) erros.ErroStatus {
	usuarioBanco, err := usuarioExternoService.usuarioExternoRepository.FindByEmail(context, email)
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
	_, err = usuarioExternoService.usuarioExternoRepository.SetTokenResetSenha(context, token, usuarioBanco.ID)
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
func (usuarioExternoService *UsuarioExternoService) AtualizarSenha(context context.Context, email string, senhaNova string) erros.ErroStatus {
	usuarioEncontrado, err := usuarioExternoService.usuarioExternoRepository.FindByEmail(context, email)
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
	_, err = usuarioExternoService.usuarioExternoRepository.UpdateSenha(context, base64.StdEncoding.EncodeToString(senhaCriptografada), usuarioEncontrado.ID)
	if err != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return erros.ErroStatus{}
}

// Função para criar uma nova usuário externo
func (usuarioExternoService *UsuarioExternoService) CreateUsuarioExterno(context context.Context, usuarioPerfilRecebido *grpc.UsuarioPerfis) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	// Busca um usuário pelo e-mail enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	_, err := usuarioExternoService.usuarioExternoRepository.FindByEmail(context, usuarioPerfilRecebido.GetUsuarioExterno().GetEmail())
	if err == nil {
		return &grpc.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("E-mail já está sendo utilizado"),
		}
	}

	// Criptografia da senha
	senhaEmBytes := []byte(usuarioPerfilRecebido.GetUsuarioExterno().GetSenha())
	senhaCriptografada, erroCriptografia := bcrypt.GenerateFromPassword(senhaEmBytes, 10)
	if erroCriptografia != nil {
		return &grpc.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroCriptografia,
		}
	}

	// Cria o objeto CreateUsuarioExternoParams gerado pelo sqlc para gravação no banco de dados
	usuarioCreate := db.CreateUsuarioExternoParams{
		Nome:      usuarioPerfilRecebido.GetUsuarioExterno().GetNome(),
		Email:     usuarioPerfilRecebido.GetUsuarioExterno().GetEmail(),
		Senha:     base64.StdEncoding.EncodeToString(senhaCriptografada),
		Uuid:      pgtype.Text{String: uuid.NewString(), Valid: true},
		IDExterno: usuarioPerfilRecebido.GetUsuarioExterno().GetIdExterno(),
		Documento: usuarioPerfilRecebido.GetUsuarioExterno().GetDocumento(),
	}

	// Cria o usuário no repositório
	usuarioExterno, err := usuarioExternoService.usuarioExternoRepository.Create(context, usuarioCreate)
	usuarioPerfilRecebido.UsuarioExterno = helpers.TUsuarioExternoToPb(usuarioExterno)
	if err != nil {
		return &grpc.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Salva a relação entre o usuário e cada perfil
	for i, perfil := range usuarioPerfilRecebido.GetPerfis() {
		perfil, err := usuarioExternoService.perfilRepository.FindByID(context, perfil.Id)
		if err != nil {
			return &grpc.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo do perfil na lista de perfis do objeto recebido
		// Como esse é o objeto retornado, os perfis retornam completos e prontos para serem usados
		usuarioPerfilRecebido.Perfis[i] = helpers.TPerfToPb(perfil)

		// Cria o objeto CreateUsuarioPerfilParams gerado pelo sqlc para gravação no banco de dados
		usuarioPerfil := db.CreateUsuarioPerfilParams{
			UsuarioExternoID: pgtype.Int4{Int32: usuarioExterno.ID, Valid: true},
			PerfilID:         perfil.ID,
			DataHora:         pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		usuarioExternoService.usuarioPerfilRepository.Create(context, usuarioPerfil)
	}

	return usuarioPerfilRecebido, erros.ErroStatus{}
}

// Função para atualizar um usuário externo existente
func (usuarioExternoService *UsuarioExternoService) UpdateUsuarioExterno(context context.Context, usuarioPerfil *pb.UsuarioPerfis, usuarioAntigo *pb.UsuarioPerfis) (*grpc.UsuarioPerfis, erros.ErroStatus) {
	usuarioRecebido := usuarioPerfil.GetUsuarioExterno()
	usuarioBanco := usuarioAntigo.GetUsuarioExterno()
	// Verifica se o e-mail foi modificado e, se sim, verifica se já existe outro registro com o mesmo e-mail
	// Em caso positivo, retorna code AlreadyExists
	if usuarioBanco.Email != usuarioRecebido.Email {
		_, err := usuarioExternoService.usuarioExternoRepository.FindByEmail(context, usuarioRecebido.Email)
		if err == nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("E-mail já está sendo utilizado"),
			}
		}
	}

	// Cria o objeto UpdateUsuarioExternoParams gerado pelo sqlc para gravação no banco de dados
	usuarioUpdate := db.UpdateUsuarioExternoParams{
		ID:        usuarioBanco.GetId(),
		Nome:      usuarioRecebido.GetNome(),
		Email:     usuarioRecebido.GetEmail(),
		Documento: usuarioRecebido.GetDocumento(),
	}

	// Salva o usuário atualizado no repositório
	usuarioAtualizado, err := usuarioExternoService.usuarioExternoRepository.Update(context, usuarioUpdate)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Deleta todas as relações antigas entre o usuário e seus perfis
	registrosExistentes, errBuscaPerfis := usuarioExternoService.usuarioPerfilRepository.FindByUsuarioExterno(context, usuarioAtualizado.ID)
	if errBuscaPerfis != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaPerfis,
		}
	}

	for i := 0; i < len(registrosExistentes); i++ {
		err := usuarioExternoService.usuarioPerfilRepository.Delete(context, registrosExistentes[i].ID)
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
		tPerfil, err := usuarioExternoService.perfilRepository.FindByID(context, perfil.Id)
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
			UsuarioExternoID: pgtype.Int4{Int32: usuarioAtualizado.ID, Valid: true},
			PerfilID:         perfil.Id,
			DataHora:         pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		usuarioExternoService.usuarioPerfilRepository.Create(context, usuarioPerfil)
	}

	usuarioPerfilRetorno := pb.UsuarioPerfis{
		UsuarioExterno: helpers.TUsuarioExternoToPb(usuarioAtualizado),
		Perfis:         perfis,
	}

	return &usuarioPerfilRetorno, erros.ErroStatus{}
}

// Função para desativar um usuário externo pelo Id
func (usuarioExternoService *UsuarioExternoService) DesativarUsuarioExterno(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Desativa o usuário externo no repositório pelo Id
	desativados, err := usuarioExternoService.usuarioExternoRepository.Desativar(context, id)

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
func (usuarioExternoService *UsuarioExternoService) RestaurarUsuarioExterno(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Desativa o usuário externo no repositório pelo Id
	restaurados, err := usuarioExternoService.usuarioExternoRepository.Restaurar(context, id)

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

// Função de login, que verifica o email, a senha do usuário e se ele está ativo ou não
// Retorna o usuário com seus perfis vinculados (UsuarioPerfis)
func (usuarioExternoService *UsuarioExternoService) Login(context context.Context, loginUsuarioExterno *pb.LoginUsuario) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuarioBanco, err := usuarioExternoService.usuarioExternoRepository.FindByEmail(context, loginUsuarioExterno.Email)
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

	perfisVinculados, erroGetPerfis := usuarioExternoService.GetPerfisVinculados(context, usuarioBanco.ID)
	if erroGetPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroGetPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfisVinculados {
		pbPerfis = append(pbPerfis, &perfisVinculados[i])
	}

	usuarioPerfil := new(pb.UsuarioPerfis)
	usuarioPerfil.UsuarioExterno = helpers.TUsuarioExternoToPb(usuarioBanco)
	usuarioPerfil.Perfis = pbPerfis

	// Validação da senha codificada do banco
	senhaHash, err := base64.StdEncoding.DecodeString(usuarioBanco.Senha)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	erroValidacao := bcrypt.CompareHashAndPassword(senhaHash, []byte(loginUsuarioExterno.Senha))
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

// Essa função serve apenas para definir a flag `SenhaAtualizada` como false
// Dessa forma o usuário consegue passar pelo middleware e acessar a API
// Essa função é chamada sempre que um usuário faz login
func (usuarioExternoService *UsuarioExternoService) SenhaAtualizada(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Salva o usuário atualizado no repositório
	atualizados, err := usuarioExternoService.usuarioExternoRepository.SetSenhaAtualizada(context, id)

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
