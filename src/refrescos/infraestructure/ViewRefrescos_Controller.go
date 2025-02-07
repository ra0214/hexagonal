package infraestructure

import (
	"actividad/src/refrescos/application"
	"encoding/json"
	"net/http"
)

type ViewRefrescosController struct {
	useCase application.ViewRefrescos
}

func NewViewRefrescosController(useCase application.ViewRefrescos) *ViewRefrescosController {
	return &ViewRefrescosController{useCase: useCase}
}

func (vp_c *ViewRefrescosController) Execute(w http.ResponseWriter, r *http.Request) {
	refrescos, err := vp_c.useCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(refrescos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
