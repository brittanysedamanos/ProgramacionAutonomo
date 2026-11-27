package productos

import (
	"errors"
	"fmt"
)

// Producto representa un articulo en el inventario
// los campos están en minuscula para encapsularlos (solo accesibles via metodos)
type Producto struct {
	sku       string
	nombre    string
	categoria string
	precio    float64
	stock     int
}

// ItemInventario es una interfaz que permite tratar productos y servicios por igual
type ItemInventario interface {
	SKU() string
	Nombre() string
	Precio() float64
	Stock() int
	ValorTotal() float64
}

// NewProducto es el constructor. Valida los datos antes de crear el producto
func NewProducto(sku, nombre, categoria string, precio float64, stockInicial int) (*Producto, error) {
	if sku == "" {
		return nil, errors.New("el SKU no puede estar vacío")
	}
	if precio < 0 {
		return nil, fmt.Errorf("el precio no puede ser negativo: %.2f", precio)
	}
	if stockInicial < 0 {
		return nil, errors.New("el stock inicial no puede ser negativo")
	}
	return &Producto{
		sku:       sku,
		nombre:    nombre,
		categoria: categoria,
		precio:    precio,
		stock:     stockInicial,
	}, nil
}

// --- GETTERS (Receptor de valor) ---
func (p Producto) SKU() string         { return p.sku }
func (p Producto) Nombre() string      { return p.nombre }
func (p Producto) Categoria() string   { return p.categoria }
func (p Producto) Precio() float64     { return p.precio }
func (p Producto) Stock() int          { return p.stock }
func (p Producto) ValorTotal() float64 { return p.precio * float64(p.stock) }

// --- SETTERS (Receptor de puntero, validan antes de modificar) ---
func (p *Producto) SetPrecio(nuevoPrecio float64) error {
	if nuevoPrecio < 0 {
		return errors.New("el precio no puede ser negativo")
	}
	p.precio = nuevoPrecio
	return nil
}

// SetStock modifica el stock controlando que no quede negativo
func (p *Producto) SetStock(nuevoStock int) error {
	if nuevoStock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	p.stock = nuevoStock
	return nil
}

// --- SLICE GLOBAL (Base de datos en memoria) ---
// Usamos un slice porque importa el orden en que se registran los productos
var inventario = []*Producto{}

// BuscarPorSKU es una función auxiliar privada (minuscula)
func buscarPorSKU(sku string) (*Producto, int, error) {
	for i, p := range inventario {
		if p.SKU() == sku {
			return p, i, nil
		}
	}
	return nil, -1, fmt.Errorf("producto con SKU %s no encontrado", sku)
}

// --- FUNCIONES EXPORTADAS ---

func AgregarProducto(sku, nombre, categoria string, precio float64, stockInicial int) error {
	// Validamos que no exista otro con el mismo SKU
	if _, _, err := buscarPorSKU(sku); err == nil {
		return fmt.Errorf("el SKU %s ya está registrado", sku)
	}
	nuevo, err := NewProducto(sku, nombre, categoria, precio, stockInicial)
	if err != nil {
		return err
	}
	inventario = append(inventario, nuevo)
	return nil
}

func BuscarProducto(sku string) (nombre, categoria string, precio float64, stock int, encontrado bool) {
	p, _, err := buscarPorSKU(sku)
	if err != nil {
		return "", "", 0, 0, false
	}
	return p.Nombre(), p.Categoria(), p.Precio(), p.Stock(), true
}

func ListarProductos() []string {
	var skus []string
	for _, p := range inventario {
		skus = append(skus, p.SKU())
	}
	return skus
}

func FiltrarPorCategoria(categoria string) []string {
	var skus []string
	for _, p := range inventario {
		if p.Categoria() == categoria {
			skus = append(skus, p.SKU())
		}
	}
	return skus
}

func ActualizarPrecio(sku string, nuevoPrecio float64) error {
	p, _, err := buscarPorSKU(sku)
	if err != nil {
		return err
	}
	return p.SetPrecio(nuevoPrecio)
}

func EliminarProducto(sku string) error {
	_, index, err := buscarPorSKU(sku)
	if err != nil {
		return err
	}
	// Removemos el elemento del slice
	inventario = append(inventario[:index], inventario[index+1:]...)
	return nil
}

// ObtenerInventario devuelve una copia del slice para que otros paquetes puedan iterarlo
func ObtenerInventario() []*Producto {
	return inventario
}
