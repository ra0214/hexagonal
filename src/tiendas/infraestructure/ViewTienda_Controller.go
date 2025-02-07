package infraestructure

import (
	"actividad/src/tiendas/application"
	"encoding/json"
	"net/http"
)

type ViewTiendaController struct {
	useCase application.ViewTienda
}

func NewViewTiendaController(useCase application.ViewTienda) *ViewTiendaController {
	return &ViewTiendaController{useCase: useCase}
}

func (vp_c *ViewTiendaController) Execute(w http.ResponseWriter, r *http.Request) {
	tiendas, err := vp_c.useCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tiendas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}