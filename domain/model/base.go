package model

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

func init() {
	// Vai validar o que estiver em valid na minha struct
	govalidator.SetFieldsRequiredByDefault(true)
}
