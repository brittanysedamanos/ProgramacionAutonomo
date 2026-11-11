package main

import (
	"fmt"
	"inventario-tienda/productos"
)

func main() {
	fmt.Println("Sistema de Inventario iniciado")
	// Ejemplo de llamada para que no quede vacío
	_ = productos.ListarProductos()
}