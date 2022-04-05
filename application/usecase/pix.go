package usecase

import (
	"errors"
	"github.com/LeonardoGuarilha/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyrepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	_, err = p.PixKeyRepository.RegisterKey(pixKey)
	if err != nil {
		return nil, errors.New("unable to create new pix key at the moment")
	}

	//if pixKey.ID == "" {
	//	return nil,
	//}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
