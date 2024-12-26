# Usa uma imagem oficial do Go
FROM golang:1.23.4

# Diretório da aplicação dentro do volume
WORKDIR /app

# Copia os arquivos do projeto atual e move para o volume
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila a aplicação
RUN go build -o main

# Define o comando de inicialização da aplicação
CMD [ "./main" ]