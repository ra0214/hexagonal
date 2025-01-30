package infraestructure

import (
	"actividad/src/refrescos/application"
	"net/http"
)

type ViewRefrescosController struct {
	useCase application.ViewRefrescos
}

func NewViewRefrescosController(useCase application.ViewRefrescos) *ViewRefrescosController {
	return &ViewRefrescosController{useCase: useCase}
}

func (vp_c *ViewRefrescosController) Execute(w http.ResponseWriter, r *http.Request) {
	vp_c.useCase.Execute()
	w.Write([]byte("Lista de Refrescos"))
}
