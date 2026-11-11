package stock

import "errors"

func RegistrarEntrada(sku string, cantidad int) error {
	return nil
}

func RegistrarSalida(sku string, cantidad int) error {
	return errors.New("no implementado")
}

func ConsultarStock(sku string) (int, error) {
	return 0, errors.New("no implementado")
}

func ProductosBajoStock(minimo int) []string {
	return []string{}
}