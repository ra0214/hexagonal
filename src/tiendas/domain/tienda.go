package domain

type ITienda interface {
	SaveTienda(nombre string, ubicacion string)
	GetAll() ([]Tienda, error)
	UpdateTienda(id int32, nombre string, ubicacion string) error
	DeleteTienda(id int32) error
}

type Tienda struct {
	ID        int32  `json:"id"`
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
}

func NewTienda(nombre string, ubicacion string) *Tienda {
	return &Tienda{ID: 1, Nombre: nombre, Ubicacion: ubicacion}
}

func (t *Tienda) GetAll() ([]Tienda, error) {
	return []Tienda{}, nil
}

func (t *Tienda) SetNombre(nombre string) {
	t.Nombre = nombre
}