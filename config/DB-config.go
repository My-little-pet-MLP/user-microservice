package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

var (
    host     string
    port     int
    user     string
    password string
    dbname   string
)

func DBConfig() (*sql.DB, error) {
	// Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar arquivo .env: %v", err)
    }

    // Assign values ​​to environment variables
    host = os.Getenv("DBHOST")
    portStr := os.Getenv("DBPORT")
    port, err = strconv.Atoi(portStr)
    if err != nil {
        log.Fatalf("Erro ao converter porta para inteiro: %v", err)
    }
    user = os.Getenv("USER")
    password = os.Getenv("PASSWORD")
    dbname = os.Getenv("DBNAME")
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Connect to PostgreSQL database
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
    }

    // Check database connection
    err = db.Ping()
    if err != nil {
        db.Close() // Close connection if Ping fails
        return nil, fmt.Errorf("erro ao pingar o banco de dados: %v", err)
    }

    fmt.Println("Conexão com o banco de dados estabelecida!")

 
    return db, nil
}

func CreateUsersTable(db *sql.DB) error {
    // Comando SQL para criar a tabela 'users'
    sqlStatement := `
         CREATE TABLE IF NOT EXISTS users (
            ID VARCHAR(255) PRIMARY KEY,
            fullname VARCHAR(255),
            imageUrl VARCHAR(255),
            email VARCHAR(255),
            phone VARCHAR(12),
            created_at TIMESTAMP
        );
    `

    // Execute o comando SQL para criar a tabela
    _, err := db.Exec(sqlStatement)
    if err != nil {
		log.Fatalf("erro ao criar tabela 'users': %v", err)
        return nil 
    }
	fmt.Println("Tabela 'users' verificada ou criada com sucesso!")

    return nil
}