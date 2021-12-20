package main

import (
	"fmt"
)

func main() {

	ejercicio1()
	fmt.Println("\n///////////////////////////////")
	ejercicio2()
	fmt.Println("\n///////////////////////////////")
	ejercicio3()

}

func ejercicio1() {

	alumn := Alumno{"Santi", "Lonzieme", 34080205, "05/09/1988"}

	alumn.details()
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) details() {
	fmt.Println("Nombre:", a.Nombre)
	fmt.Println("Apellido:", a.Apellido)
	fmt.Println("DNI:", a.DNI)
	fmt.Println("Fecha:", a.Fecha)
}

func ejercicio2() {

	var matrizValues = [9]float64{4, 6, 8, 2, 8, 2, 4, 5, 6}

	matr := Matriz{alto: 3, ancho: 3, cuadratica: true, valorMax: 8}

	matr.setMatriz(matrizValues)
	matr.print()

}

type Matriz struct {
	valoresMatriz [9]float64
	alto          float64
	ancho         float64
	cuadratica    bool
	valorMax      float64
}

func (m *Matriz) setMatriz(v [9]float64) {
	m.valoresMatriz = v
}

func (m Matriz) print() {

	for i := range m.valoresMatriz {

		if i == 2 || i == 5 {
			fmt.Println(m.valoresMatriz[i])
		} else {
			fmt.Print(m.valoresMatriz[i])
		}
	}
}

func ejercicio3() {

	producto1 := producto{"Bicicleta", 10000, "Grande"}
	producto2 := producto{"Zapatillas", 1500, "Mediano"}
	producto3 := producto{"Cuaderno", 100, "Pequeño"}
	nuevaTienda := tienda{}
	nuevaTienda.Agregar(producto1)
	nuevaTienda.Agregar(producto2)
	nuevaTienda.Agregar(producto3)

	fmt.Println(nuevaTienda.lista)
	nuevaTienda.Total()

}

type tienda struct {
	lista []producto
}

type producto struct {
	Nombre string
	Precio float64
	Tipo   string
}

type Producto interface {
	CalcularCosto()
}

type Ecommerce interface {
	Total()
	Agregar()
}

func (p *producto) nuevoProducto(tipo string, nombre string, precio float64) producto {

	producto := producto{nombre, precio, tipo}

	return producto
}

func CalcularCosto(p producto) (costoAdi float64) {

	switch p.Tipo {
	case "Pequeño":
		costoAdi = 0
		fmt.Println("El costo adicional es 0")
	case "Mediano":
		costoAdi = p.Precio / 100 * 3
		fmt.Println("El costo adicional del producto mediano es", costoAdi)
	case "Grande":
		costoAdi = p.Precio/100*6 + 2500
		fmt.Println("El costo adicional del producto grande es", costoAdi)
	}

	return
}

func (t *tienda) Agregar(producto producto) (NuevaTienda tienda) {

	t.lista = append(t.lista, producto)

	return *t

}

func (t *tienda) Total() (precioFinal float64) {

	for _, producto := range t.lista {
		precioFinal = precioFinal + CalcularCosto(producto) + producto.Precio
	}

	fmt.Println("El total a pagar de la lista es", precioFinal)
	return
}
