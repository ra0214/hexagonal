package main

import (
	tiendaInfra "actividad/src/tiendas/infraestructure"
	perfumeInfra "actividad/src/refrescos/infraestructure"
)

func main() {
	tiendaInfra.Iniciar()

	perfumeInfra.Init()
}
