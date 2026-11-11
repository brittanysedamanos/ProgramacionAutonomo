Inventario de Tienda
Sistema de gestión de inventario para una tienda pequeña desarrollado en Go como parte del Aprendizaje Autónomo 1 de la Unidad 1. 
El proyecto busca que el dueño o encargado pueda llevar un control ordenado de sus productos sin depender de un cuaderno o de una hoja de Excel que con el tiempo se vuelve difícil de mantener. 
La organización en paquetes y funciones permite que el sistema crezca por etapas sin abrumar al usuario con una herramienta compleja desde el primer día.

Objetivo
Diseñar la estructura de un sistema de inventario organizado en paquetes y funciones de Go que permita registrar productos y consultarlos y modificarlos y eliminarlos. 
También busca llevar el control de entradas y salidas de stock y generar reportes básicos que ayuden a tomar decisiones. 
En esta primera etapa solo se declaran las funciones y se organizan los paquetes. 
La lógica completa se implementará en etapas posteriores del curso.

Alcance
El sistema puede registrar productos con su SKU su nombre su categoría su precio y su stock inicial. 
Puede consultar un producto por su SKU y puede listar todos los productos registrados. 
Puede filtrar por categoría y puede modificar el precio de un producto y puede eliminarlo del inventario.

En cuanto al stock el sistema puede registrar entradas cuando llega mercadería y puede registrar salidas cuando se vende algo. 
Puede consultar el stock actual de cualquier producto y puede detectar cuáles están por debajo de un mínimo definido.

Por el lado de los reportes el sistema puede calcular el valor total del inventario y puede generar un resumen de cuántos productos hay por categoría. 
También puede mostrar un listado legible de los productos que están bajos de stock.

Todavía no incluye una interfaz web. 
Tampoco maneja usuarios ni login ni roles. 
No se conecta a una base de datos porque en esta etapa los datos se manejan en memoria. 
No procesa pagos reales ni facturación. 
Tampoco genera archivos PDF ni exportaciones porque los listados son en texto.

Módulos y paquetes
El sistema se divide en tres módulos y cada uno se implementa como un paquete de Go. 
El módulo de productos corresponde al paquete productos y su archivo es productos.go dentro de la carpeta productos. 
El módulo de stock corresponde al paquete stock y su archivo es stock.go dentro de la carpeta stock. 
El módulo de reportes corresponde al paquete reportes y su archivo es reportes.go dentro de la carpeta reportes. 
Los nombres van en minúsculas y sin guiones bajos como recomienda la guía de estilo de Go. 
El paquete main es el único que no corresponde a un módulo del sistema: su función es arrancar el programa y llamar a los demás paquetes. 
La lógica no se concentra en main.

    
Paquetes
En esta etapa no se usan paquetes de terceros. 
Todo se resuelve con la librería estándar de Go. 
Se usan errors y fmt y sort y strings. 
Si en etapas posteriores se incorpora alguno se documentará su uso y su justificación.

Cómo compilar y ejecutar
Para comprobar que el proyecto compila hay que ubicarse dentro de la carpeta del repositorio y ejecutar go build ./... 
Si no muestra ningún error significa que la estructura de paquetes y las firmas de las funciones están bien. 
Para ejecutar el programa se usa go run main.go.

Estado del proyecto
Etapa 1 completada: estructura de paquetes y funciones declaradas. 
En la siguiente etapa se incorporarán structs y encapsulación y manejo de errores e interfaces. 
Al final del curso se convertirá en una aplicación web con al menos ocho servicios.
