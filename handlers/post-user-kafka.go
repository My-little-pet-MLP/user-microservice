package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/my-little-pet/user-microservice/config"
	"github.com/my-little-pet/user-microservice/models"
	
)

func PostUserKafka( w http.ResponseWriter,r *http.Request) {

	// instanciando a config do kafka
	producer := config.KafkaConfigProducer()
	defer producer.Close()

	// garantindo que o metodo é post
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	var user models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}
	
	// criando um userfull
	var userfull models.User
	userfull.ID = uuid.NewString()
	userfull.Fullname = user.Fullname
	userfull.ImageUrl = user.ImageUrl
	userfull.Email = user.Email
	userfull.Phone = user.Phone
	userfull.CreatedAt = time.Now()
	
	userData, err := json.Marshal(userfull)
	if err != nil {
		http.Error(w, "Erro ao codificar o usuário", http.StatusInternalServerError)
		return
	}
	log.Println(user)
	message := &sarama.ProducerMessage{
		Topic: "users",
		Value: sarama.ByteEncoder(userData),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		http.Error(w, "Erro ao enviar mensagem para o Kafka", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Mensagem enviada para o tópico users, na partição %d, no offset %d \n", partition, offset)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário recebido e mensagem publicada no Kafka"))
}