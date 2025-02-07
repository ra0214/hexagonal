package main

import (
	"fmt"
	tiendaInfra "actividad/src/tiendas/infraestructure"
	refrescosInfra "actividad/src/refrescos/infraestructure"
)

func main() {
	fmt.Println("Inicializando tienda...")
	go tiendaInfra.Init()

	fmt.Println("Inicializando refrescos...")
	go refrescosInfra.Init()

	select {}
}
