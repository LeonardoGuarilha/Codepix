package main

import (
	"github.com/LeonardoGuarilha/codepix-go/application/grpc"
	"github.com/LeonardoGuarilha/codepix-go/infra/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051) //50051 porta padrão do grpc
}
