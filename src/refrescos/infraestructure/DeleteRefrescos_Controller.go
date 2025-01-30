package infraestructure

import (
	"actividad/src/refrescos/application"
	"fmt"
	"net/http"
	"strconv"
)

type DeleteRefrescosController struct {
	useCase application.DeleteRefrescos
}

func NewDeleteRefrescosController(useCase application.DeleteRefrescos) *DeleteRefrescosController {
	return &DeleteRefrescosController{useCase: useCase}
}

func (dp_c *DeleteRefrescosController) Execute(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de bebida inválido", http.StatusBadRequest)
		return
	}

	err = dp_c.useCase.Execute(int32(id))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar el refresco: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Refresco eliminado correctamente"))
}
