package productos

import "errors"

func AgregarProducto(sku, nombre, categoria string, precio float64, stockInicial int) error {
	return nil
}

func BuscarProducto(sku string) (nombre, categoria string, precio float64, stock int, encontrado bool) {
	return "", "", 0, 0, false
}

func ListarProductos() []string {
	return []string{}
}

func FiltrarPorCategoria(categoria string) []string {
	return []string{}
}

func ActualizarPrecio(sku string, nuevoPrecio float64) error {
	return errors.New("no implementado")
}

func EliminarProducto(sku string) error {
	return errors.New("no implementado")
}