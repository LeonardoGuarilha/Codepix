package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"), // Servidor de conexão com o kafka
	}

	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic:     &topic,
			Partition: ckafka.PartitionAny, // fica automático para o kafka colocar em qualquer partição
		},
		Value: []byte(msg), // a mensagem
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

// Fica rodando em background
// Sempre que uma mensagem cair no chanel, essa mensagem vai pegar
func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch event := e.(type) {
		case *ckafka.Message:
			if event.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", event.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", event.TopicPartition)
			}

		}
	}
}
