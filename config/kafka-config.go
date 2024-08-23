package config

import (
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/my-little-pet/user-microservice/utils"
)
var (
	kafkaaddress string
)
func KafkaConfigProducer() (sarama.SyncProducer){
		// Verifica se a variável de ambiente KAFKAADDRESS está definida
		utils.CheckEnvVar("KAFKAADDRESS")

		// Obtém o valor de KAFKAADDRESS do ambiente
		kafkaaddress = os.Getenv("KAFKAADDRESS")
		log.Println("Kafka Address:", kafkaaddress)
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{kafkaaddress}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	return producer
}

func KafkaConfigConsumer() (sarama.Consumer){
	// Verifica se a variável de ambiente KAFKAADDRESS está definida
	utils.CheckEnvVar("KAFKAADDRESS")

	// Obtém o valor de KAFKAADDRESS do ambiente
	kafkaaddress = os.Getenv("KAFKAADDRESS")
	log.Println("Kafka Address:", kafkaaddress)

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{kafkaaddress}

	fmt.Println("Iniciando consumidor Kafka...")

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Erro ao criar consumidor: %v", err)
	}
	return consumer
}