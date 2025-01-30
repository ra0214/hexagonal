package application

import "actividad/src/refrescos/domain"

type ViewRefrescos struct {
	db domain.IRefrescos
}

func NewViewRefrescos(db domain.IRefrescos) *ViewRefrescos {
	return &ViewRefrescos{db: db}
}

func (vp *ViewRefrescos) Execute()  {
	vp.db.GetAll()
}