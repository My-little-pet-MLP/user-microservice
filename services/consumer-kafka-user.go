package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/my-little-pet/user-microservice/config"
	"github.com/my-little-pet/user-microservice/models"

	"github.com/go-playground/validator/v10"
)

func Executeconsumer() {
	consumer := config.KafkaConfigConsumer()
	defer func() {
		fmt.Println("Fechando consumidor...")
		if err := consumer.Close(); err != nil {
			log.Fatalf("Erro ao fechar consumidor: %v", err)
		}
	}()

	fmt.Println("Consumidor criado com sucesso. Iniciando consumo de partição...")

	partitionConsumer, err := consumer.ConsumePartition("users", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Erro ao consumir partição: %v", err)
	}

	fmt.Println("Consumindo partição com sucesso.")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	messages := partitionConsumer.Messages()

	for {
		select {
		case msg := <-sigCh:
			fmt.Printf("Saindo da aplicação, código recebido %v", msg)
			return
		case msg := <-messages:
			var user models.User
			if err := json.Unmarshal(msg.Value, &user); err != nil {
				log.Printf("Erro ao desserializar mensagem JSON: %v\n", err)
			}
			fmt.Println(user)
			
			db, err := config.DBConfig()
			if err != nil {
				log.Fatal("Erro ao conectar ao banco de dados")
				return
			}
	

			validate := validator.New(validator.WithRequiredStructEnabled())
			errvalidade := validate.Struct(user)
			if errvalidade != nil {
				fmt.Println("Erro ao validar o usuario: " + errvalidade.Error())
				continue
			} else {
				// Verificando se o usuário já existe pelo email
				userEmailExists,err := GetByEmailUser(user.Email)
				if err != nil {
					fmt.Println("Erro ao buscar o user!");
					continue
				}
				if userEmailExists != nil {
					fmt.Println("Usuário com email já cadastrado: " + user.Email)
					continue
				}
				sqlStatement := `
		  INSERT INTO users (id, fullname, imageUrl, email, phone,created_at)
		  VALUES ($1, $2, $3, $4, $5, $6)`

				_, errdb := db.Exec(sqlStatement,user.ID, user.Fullname, user.ImageUrl, user.Email, user.Phone, user.CreatedAt)
				if errdb != nil {
					fmt.Println("Erro ao inserir no banco de dados: " + errdb.Error())
					continue
				}
				log.Println("Usuario salvo com sucesso!")

			}

		}
	}
}