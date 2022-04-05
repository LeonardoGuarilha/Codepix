package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"` // FK com bank
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"` // Posso ter mais de 1 pix key por conta
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	accont := Account{
		OwnerName: ownerName,
		BankID:    bank.ID,
		Number:    number,
	}

	accont.ID = uuid.NewV4().String()
	accont.CreatedAt = time.Now()
	//accont.UpdatedAt = time.Now()

	err := accont.isValid()
	if err != nil {
		return nil, err
	}

	return &accont, nil
}
