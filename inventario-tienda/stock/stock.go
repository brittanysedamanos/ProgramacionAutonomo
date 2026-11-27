package stock

import (
	"errors"
	"inventario-tienda/productos"
)

func RegistrarEntrada(sku string, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad a ingresar debe ser mayor a cero")
	}

	// Buscamos el producto para actualizarlo
	_, categoria, _, stockActual, encontrado := productos.BuscarProducto(sku)
	if !encontrado {
		return errors.New("el producto no existe")
	}
	_ = categoria // No lo usamos aqui pero lo obtenemos de la búsqueda

	// Usamos el setter de productos para modificar
	return actualizarStockInterno(sku, stockActual+cantidad)
}

func RegistrarSalida(sku string, cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad a retirar debe ser mayor a cero")
	}

	_, _, _, stockActual, encontrado := productos.BuscarProducto(sku)
	if !encontrado {
		return errors.New("el producto no existe")
	}

	if stockActual < cantidad {
		return errors.New("no hay suficiente stock para realizar la salida")
	}

	return actualizarStockInterno(sku, stockActual-cantidad)
}

func ConsultarStock(sku string) (int, error) {
	_, _, _, stockActual, encontrado := productos.BuscarProducto(sku)
	if !encontrado {
		return 0, errors.New("el producto no existe")
	}
	return stockActual, nil
}

func ProductosBajoStock(minimo int) []string {
	var skus []string
	for _, p := range productos.ObtenerInventario() {
		if p.Stock() <= minimo {
			skus = append(skus, p.SKU())
		}
	}
	return skus
}

// Función auxiliar para no repetir código al modificar el stock
func actualizarStockInterno(sku string, nuevoStock int) error {
	// Como el paquete productos no tiene un SetStock público directo por SKU
	// usamos ActualizarPrecio como ejemplo pero necesitamos modificar stock
	// para hacerlo simple en el paquete productos podríamos agregar una función
	// pero mjr usamos los métodos disponibles.

	// Lo ideal sería que productos tenga un SetStockPorSKU.
	// lo simulamos:
	for _, p := range productos.ObtenerInventario() {
		if p.SKU() == sku {
			return p.SetStock(nuevoStock) // Usamos el setter del struct
		}
	}
	return errors.New("error interno: no se encontró el producto")
}
