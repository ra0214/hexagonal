package infraestructure

import (
	"actividad/src/refrescos/application"
	"encoding/json"
	"net/http"
)

type CreateRefrescosController struct {
	useCase application.CreateRefrescos
}

func NewCreateRefrescosController(useCase application.CreateRefrescos) *CreateRefrescosController {
	return &CreateRefrescosController{useCase: useCase}
}

type RequestBody struct {
	Marca  string  `json:"marca"`
	Precio float32 `json:"precio"`
}

func (cp_c *CreateRefrescosController) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error al leer el JSON", http.StatusBadRequest)
		return
	}

	cp_c.useCase.Execute(body.Marca, body.Precio)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Refresco agregado correctamente"})
}
