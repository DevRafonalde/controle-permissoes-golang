FROM golang:1.23.2

# Criar diretório de trabalho fora do GOPATH
WORKDIR /app

# Copiar go.mod e go.sum primeiro para aproveitar o cache de build
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o restante do código
COPY . .

# Construir o aplicativo
RUN go build -o main .

# Comando padrão para rodar o binário
CMD ["./main"]
