package application

import (
	"actividad/src/refrescos/domain"
)

type EditRefrescos struct {
	db domain.IRefrescos
}

func NewEditRefrescos(db domain.IRefrescos) *EditRefrescos {
	return &EditRefrescos{db: db}
}

func (ep *EditRefrescos) Execute(id int32, marca string, precio float32) error {
	return ep.db.UpdateRefrescos(id, marca, precio)
}