package infraestructure

import (
	"actividad/src/refrescos/application"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type EditRefrescosController struct {
	useCase application.EditRefrescos
}

func NewEditRefrescosController(useCase application.EditRefrescos) *EditRefrescosController {
	return &EditRefrescosController{useCase: useCase}
}

func (ep_c *EditRefrescosController) Execute(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de bebida inválido", http.StatusBadRequest)
		return
	}

	var p struct {
		Marca  string  `json:"marca"`
		Precio float32 `json:"precio"`
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}

	err = ep_c.useCase.Execute(int32(id), p.Marca, p.Precio)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar el refresco: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Refresco actualizado correctamente"})
}