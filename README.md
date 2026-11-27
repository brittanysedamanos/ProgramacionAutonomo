Inventario de Tienda

Es un sistema de gestión de inventario para una tienda pequeña desarrollado en Go. 
El proyecto busca que la dueña o encargada pueda llevar un control ordenado de sus productos sin depender de un cuaderno 
o de una hoja de Excel que con el tiempo se vuelve difícil de mantener. La organización en paquetes permite que el 
sistema crezca por etapas sin abrumar al usuario con una herramienta compleja desde el primer día, teniendo en cuenta 
que el público objetivo son personas mayores que necesitan interfaces y mensajes claros.


Objetivo

Diseñar e implementar un sistema de inventario organizado en paquetes de Go que permita registrar productos, 
consultarlos, modificarlos y eliminarlos. También busca llevar el control de entradas y salidas de stock y generar 
reportes básicos que ayuden a tomar decisiones. En esta segunda etapa, el sistema ya cuenta con la lógica implementada 
usando structs, encapsulación y manejo de errores.


Alcance

El sistema puede registrar productos con su SKU, nombre, categoría, precio y stock inicial. Puede consultar un producto 
por su SKU y puede listar todos los productos registrados. Puede filtrar por categoría, modificar el precio y eliminar 
productos.


En cuanto al stock, el sistema puede registrar entradas cuando llega mercadería y salidas cuando se vende. 
Puede consultar el stock actual y detectar cuáles están por debajo de un mínimo definido.


Por el lado de los reportes, el sistema calcula el valor total del inventario, genera un resumen por categoría y muestra
 los productos bajos de stock.


Todavía no incluye una interfaz web, no maneja usuarios ni login, y no se conecta a una base de datos porque los datos 
se manejan en memoria.


Módulos y paquetes

El sistema se divide en tres módulos y cada uno se implementa como un paquete de Go:

El módulo de productos corresponde al paquete productos.

El módulo de stock corresponde al paquete stock.

El módulo de reportes corresponde al paquete reportes.

El paquete main es el único que no corresponde a un módulo del sistema: su función es arrancar el programa, 
probar las funcionalidades y llamar a los demás paquetes. La lógica de negocio no se concentra en main.


Paquetes de terceros

En esta etapa no se usan paquetes de terceros. Todo se resuelve con la librería estándar de Go. 
Se usan errors, fmt y las estructuras básicas del lenguaje. Si en etapas posteriores se incorpora alguno, 
se documentará su uso y su justificación.


Decisiones de diseño 

Para esta etapa se tomaron las siguientes decisiones, basadas en las buenas prácticas de Go:

Entidades y Structs

Se creó el struct Producto para modelar el artículo físico. 
Los campos están en minúscula para que sean privados.

Receptores

Los getters usan receptor de valor (p Producto) porque solo leen información. 
Los setters usan receptor de puntero (p *Producto) porque necesitan modificar la instancia original y no una copia.

Estructuras de datos

Se utilizó un slice ([]*Producto) para almacenar el inventario porque el orden de inserción 
es importante para el usuario. Si el sistema creciera mucho, se podría migrar a un map para búsquedas más rápidas, 
pero para este alcance el slice es óptimo.

Interfaz

Se creó la interfaz ItemInventario que permite aplicar polimorfismo en los reportes. 
El cálculo del valor total funciona con cualquier entidad que cumpla con esa interfaz.


Aplicación de conceptos teóricos

Encapsulación

Se aplicó ocultando los campos del struct Producto. 
Para modificar el estado se usan métodos setter que validan la información. 
Siguiendo la teoría, el setter usa un puntero para que el cambio persista en la estructura original y no en una copia.

Manejo de errores 

Se utilizó errors.New para validaciones fijas (como evitar SKU vacíos) y fmt.Errorf cuando era 
necesario incluir el valor erróneo en el mensaje (como un precio negativo). 
El programa principal (main.go) revisa cada error devuelto (if err != nil) evitando que el sistema falle abruptamente.

Interfaces

Al usar la interfaz ItemInventario, el sistema puede tratar de la misma forma a los productos actuales 
y a futuros servicios sin tener que reescribir la lógica de reportes.



