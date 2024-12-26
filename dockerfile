# Usa uma imagem oficial do Go
# Tomar bastante cuidado com a versão que esta no dockerfile e no go.mod, pois gera conflito e erro na hora do build
FROM golang:1.21

# OBS -> O container não roda na versão 1.23.4 do Go, investigando o por isto acontece

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