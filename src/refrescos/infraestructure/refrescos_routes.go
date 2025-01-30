package infraestructure

import (
	"actividad/src/refrescos/application"
	"actividad/src/refrescos/domain"
	"net/http"
)

func SetupRouter(repo domain.IRefrescos) {
	createRefrescos := application.NewCreateRefrescos(repo)
	createRefrescosController := NewCreateRefrescosController(*createRefrescos)

	viewRefrescos := application.NewViewRefrescos(repo)
	viewRefrescosController := NewViewRefrescosController(*viewRefrescos)

	editRefrescosUseCase := application.NewEditRefrescos(repo)
	editRefrescosController := NewEditRefrescosController(*editRefrescosUseCase)

	deleteRefrescosUseCase := application.NewDeleteRefrescos(repo)
	deleteRefrescosController := NewDeleteRefrescosController(*deleteRefrescosUseCase)

	http.HandleFunc("/refrescos", createRefrescosController.Execute)
	http.HandleFunc("/refresco", viewRefrescosController.Execute)
	http.HandleFunc("/editRefrescos", editRefrescosController.Execute)
	http.HandleFunc("/deleteRefrescos", deleteRefrescosController.Execute)

	http.ListenAndServe(":8080", nil)
}
