package domain

type IRefrescos interface {
	SaveRefrescos(marca string, precio float32)
	GetAll() ([]Refrescos, error)
	UpdateRefrescos(id int32, marca string, precio float32) error
	DeleteRefrescos(id int32) error
}

type Refrescos struct {
	ID     int32   `json:"id"`
	Marca  string  `json:"marca"`
	Precio float32 `json:"precio"`
}

func NewRefrescos(marca string, precio float32) *Refrescos {
	return &Refrescos{ID:1, Marca: marca, Precio: precio}
}

func (p *Refrescos) GetAll() ([]Refrescos, error) {
	return []Refrescos{}, nil
}

func (p *Refrescos) SetMarca(marca string) {
	p.Marca = marca
}