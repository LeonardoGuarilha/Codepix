/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/LeonardoGuarilha/codepix-go/application/kafka"
	"github.com/LeonardoGuarilha/codepix-go/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"

	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using apache kafka",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Produzindo mensagem")

		// Cria a deliveryChan
		deliveryChan := make(chan ckafka.Event)
		database := db.ConnectDB(os.Getenv("env"))

		producer := kafka.NewKafkaProducer()
		// Todas as mensasgens vão ir para esse deliveryChan
		//kafka.Publish("Olá consumer", "teste", producer, deliveryChan)

		// Vai rodar em uma thread separada e não vai travar o processamento do kaftaProcessor abaixo
		// Roda como se fosse de forma asyncrona
		go kafka.DeliveryReport(deliveryChan)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
