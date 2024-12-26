package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main () {
	// Buscas as variáveis de ambiente do sistema
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Monta a query de conexão para utilizar com a lib padrão de database
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	
	// Cria a conexão com o banco de dados
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}
	defer db.Close()

	// Testa a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados bem sucedida!")

	// Teste de rota inicial
	// Na declaração do handler de requests precisa fazer um ponteiro para a api interna de HTTP.
	// TODO: Tirar dúvida por que puxa um ponteiro no request
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request){
		// Na hora de enviar o valor com o response.Write precisa estipular o tipo de dados.
		// TODO: tirar dúvida por que usam byte
		response.Write([]byte("API GoLang rodando com postgresSQL"))
	})

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}