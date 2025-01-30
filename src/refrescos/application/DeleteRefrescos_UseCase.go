package application

import "actividad/src/refrescos/domain"

type DeleteRefrescos struct {
	db domain.IRefrescos
}

func NewDeleteRefrescos(db domain.IRefrescos) *DeleteRefrescos {
	return &DeleteRefrescos{db: db}
}

func (dp *DeleteRefrescos) Execute(id int32) error {
	return dp.db.DeleteRefrescos(id)
}
