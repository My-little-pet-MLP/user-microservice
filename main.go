package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/my-little-pet/user-microservice/config"
	"github.com/my-little-pet/user-microservice/handlers"
	service "github.com/my-little-pet/user-microservice/services"
)


func main() {

	
	// Configuração do banco de dados
	db, err := config.DBConfig()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	db.Close()
	r := mux.NewRouter()
		// Definir rotas usando gorilla/mux
	r.HandleFunc("/users", handlers.PostUserKafka).Methods("POST")
	r.HandleFunc("/hearth", handlers.HearthHandlerfunc).Methods("GET")
	r.HandleFunc("/users/id={id}", handlers.GetByIdUserHandler).Methods("GET")
	r.HandleFunc("/users/email={email}", handlers.GetByEmailUserHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
		// Configurar para ouvir sinais do sistema operacional
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	// Iniciar o consumidor Kafka em uma goroutine
	go func() {
		service.Executeconsumer()
	}()

	go func() {
		fmt.Println("Servidor HTTP rodando na porta 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao inicar o servidor HTTP: %v", err)
		}
	}()
	
		// Aguardar sinal para encerramento
		sig := <-sigCh
		fmt.Printf("Recebido sinal %v. Encerrando...\n", sig)
	
		// Encerrar servidor HTTP
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Erro ao encerrar o servidor HTTP: %v", err)
		}
}