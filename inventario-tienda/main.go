package main

import (
	"fmt"
	"inventario-tienda/productos"
	"inventario-tienda/reportes"
	"inventario-tienda/stock"
)

func main() {
	fmt.Println("=== SISTEMA DE INVENTARIO ===")

	// 1. Agregar productos (manejo de errores)
	fmt.Println("\n-- Agregando productos --")
	err := productos.AgregarProducto("A001", "Cuaderno", "Papelería", 2.50, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = productos.AgregarProducto("A002", "Lápiz", "Papelería", 1.00, 5)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Intentamos agregar uno con SKU repetido para probar el error
	err = productos.AgregarProducto("A001", "Borrador", "Papelería", 0.50, 20)
	if err != nil {
		fmt.Println("Error esperado:", err) // Debe mostrar que el SKU ya existe
	}

	// 2. Listar
	fmt.Println("\n-- Lista de SKUs --")
	fmt.Println(productos.ListarProductos())

	// 3. Modificar stock
	fmt.Println("\n-- Movimientos de stock --")
	err = stock.RegistrarEntrada("A001", 5)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = stock.RegistrarSalida("A002", 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Error esperado: intentar sacar mas de lo que hay
	err = stock.RegistrarSalida("A002", 100)
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	// 4. Consultar el stock
	s, err := stock.ConsultarStock("A001")
	if err == nil {
		fmt.Printf("Stock de A001: %d\n", s)
	}

	// 5. Reportes
	fmt.Println("\n-- Reportes --")
	fmt.Printf("Valor total del inventario: $%.2f\n", reportes.ValorTotalInventario())
	fmt.Printf("Productos por categoría: %v\n", reportes.ResumenPorCategoria())

	// 6. Actualizar el precio con la validación
	err = productos.ActualizarPrecio("A001", -5.00)
	if err != nil {
		fmt.Println("Error esperado (precio negativo):", err)
	}
}
