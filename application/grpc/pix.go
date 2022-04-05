package grpc

import (
	"context"
	"github.com/LeonardoGuarilha/codepix-go/application/grpc/pb"
	"github.com/LeonardoGuarilha/codepix-go/application/usecase"
)

type PixGrpcService struct {
	PixUseCase                       usecase.PixUseCase
	pb.UnimplementedPixServiceServer // Tem que implementar por causa do grpc
}

// Implemento o método da interface que foi gerada no pixkey_grpc.pb.go
// Com isso, implemento os métodos da interface do grpc para que a comunicação
// seja feita com grpc
func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "Not Created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "Created",
	}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindKey(in.Key, in.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.ID,
		Kind: pixKey.Kind,
		Key:  pixKey.Key,
		Account: &pb.Account{
			AccountId:     pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

func NewPixGrpcService(useCase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: useCase,
	}
}
