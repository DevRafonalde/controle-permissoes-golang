package helpers

import (
	"si-admin/app/model/grpc"
	db "si-admin/app/model/repositories/sqlc/repositoryIMPL"
)

// Conversão de TUsuario para pb.Usuario
func TUsuarioInternoToPb(usuario db.TUsuariosInterno) *grpc.UsuarioInterno {
	return &grpc.UsuarioInterno{
		Id:              usuario.ID,
		Nome:            usuario.Nome,
		Email:           usuario.Email,
		Senha:           usuario.Senha,
		Ativo:           usuario.Ativo.Bool,
		TokenResetSenha: usuario.TokenResetSenha.String,
		AtualizadoEm:    usuario.AtualizadoEm.Time.Format("02/01/2006"),
		SenhaAtualizada: usuario.SenhaAtualizada.Bool,
	}
}

// Conversão de TPerfi para pb.Perfil
func TPerfToPb(perfil db.TPerfi) *grpc.Perfil {
	return &grpc.Perfil{
		Id:           perfil.ID,
		Nome:         perfil.Nome,
		Descricao:    perfil.Descricao,
		Ativo:        perfil.Ativo.Bool,
		AtualizadoEm: perfil.AtualizadoEm.Time.Format("02/01/2006"),
	}
}

func TPermissaoToPb(permissao db.TPermisso) *grpc.Permissao {
	return &grpc.Permissao{
		Id:           permissao.ID,
		Nome:         permissao.Nome,
		Descricao:    permissao.Descricao,
		Ativo:        permissao.Ativo.Bool,
		AtualizadoEm: permissao.AtualizadoEm.Time.Format("02/01/2006"),
	}
}

func TUsuarioExternoToPb(usuarioExterno db.TUsuariosExterno) *grpc.UsuarioExterno {
	return &grpc.UsuarioExterno{
		Id:              usuarioExterno.ID,
		Uuid:            usuarioExterno.Uuid.String,
		IdExterno:       usuarioExterno.IDExterno,
		Nome:            usuarioExterno.Nome,
		Email:           usuarioExterno.Email,
		Senha:           usuarioExterno.Senha,
		Documento:       usuarioExterno.Documento,
		CodReserva:      usuarioExterno.CodReserva.String,
		Ativo:           usuarioExterno.Ativo.Bool,
		TokenResetSenha: usuarioExterno.TokenResetSenha.String,
		AtualizadoEm:    usuarioExterno.AtualizadoEm.Time.Format("02/01/2006"),
		SenhaAtualizada: usuarioExterno.SenhaAtualizada.Bool,
	}
}
