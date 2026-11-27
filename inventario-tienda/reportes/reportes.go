package reportes

import (
	"inventario-tienda/productos"
)

// ValorTotalInventario suma el valor de todos los items usando la interfaz
func ValorTotalInventario() float64 {
	total := 0.0
	for _, p := range productos.ObtenerInventario() {
		// p es de tipo *Producto pero lo tratamos como ItemInventario
		var item productos.ItemInventario = p
		total += item.ValorTotal()
	}
	return total
}

func ResumenPorCategoria() map[string]int {
	resumen := make(map[string]int)
	for _, p := range productos.ObtenerInventario() {
		resumen[p.Categoria()]++
	}
	return resumen
}

func ListadoBajoStock(minimo int) []string {
	var listado []string
	for _, p := range productos.ObtenerInventario() {
		if p.Stock() <= minimo {
			// Usamos fmt implícito o simplemente devolvemos el SKU
			listado = append(listado, p.SKU())
		}
	}
	return listado
}
