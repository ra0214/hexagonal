package domain

type IRefrescos interface {
	SaveRefrescos(marca string, precio float32)
	GetAll()
	UpdateRefrescos(id int32, marca string, precio float32) error
	DeleteRefrescos(id int32) error
}

type refrescos struct {
	ID     int32   `json:"id"`
	Marca  string  `json:"marca"`
	Precio float32 `json:"precio"`
}

func NewRefrescos(marca string, precio float32) *refrescos {
	return &refrescos{ID:1, Marca: marca, Precio: precio}
}

func (p *refrescos) GetAll() ([]refrescos, error) {
	return []refrescos{}, nil
}

func (p *refrescos) SetMarca(marca string) {
	p.Marca = marca
}