package erros

import "google.golang.org/grpc/codes"

// Esse tipo é usado nos SERVICES, pois eles sabem exatamente qual erro aconteceu e já retornam para o controller o erro e o código corretos
type ErroStatus struct {
	Status codes.Code // Código de erro do gRPC
	Erro   error      // Mensagem de erro
}
