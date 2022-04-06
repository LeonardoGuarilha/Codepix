package factory

import (
	"github.com/LeonardoGuarilha/codepix-go/application/usecase"
	"github.com/LeonardoGuarilha/codepix-go/infra/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixKeyRepository:      pixRepository,
	}

	return transactionUseCase
}
