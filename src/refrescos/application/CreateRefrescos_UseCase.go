package application

import (
	"actividad/src/refrescos/domain"
)

type CreateRefrescos struct {
	db domain.IRefrescos
}

func NewCreateRefrescos(db domain.IRefrescos) *CreateRefrescos {
	return &CreateRefrescos{db: db}
}

func (cp *CreateRefrescos) Execute(marca string, precio float32) {
	cp.db.SaveRefrescos(marca, precio)
}
