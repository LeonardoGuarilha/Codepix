/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/LeonardoGuarilha/codepix-go/application/grpc"
	"github.com/LeonardoGuarilha/codepix-go/application/kafka"
	"github.com/LeonardoGuarilha/codepix-go/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"

	"github.com/spf13/cobra"
)

var (
	gRPCPortNumber int
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and Kafka Consumers",

	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		// inicia o servidor gRPC em uma outra thread
		go grpc.StartGrpcServer(database, gRPCPortNumber)

		// Cria a deliveryChan
		deliveryChan := make(chan ckafka.Event)
		//database := db.ConnectDB(os.Getenv("env"))

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
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRPCPortNumber, "grpc-port", "p", 50051, "gRPC Port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
