package main

import (
	"fmt"
	"time"
)

func main() {

	prod1 := Productos{"notebook", 24, 2}
	prod2 := Productos{"bicileta", 10000, 1}
	prod3 := Productos{"guitarra", 400, 1}
	productos := []Productos{prod1, prod2, prod3}

	serv1 := Servicio{"lavanderia", 500, 60}
	serv2 := Servicio{"seguridad", 300, 120}
	serv3 := Servicio{"envios", 700, 160}
	servicios := []Servicio{serv1, serv2, serv3}

	man1 := Mantenimiento{"Pared", 5000}
	man2 := Mantenimiento{"Reposteria", 15000}
	man3 := Mantenimiento{"Techo", 5000}
	mantenimiento := []Mantenimiento{man1, man2, man3}

	a := <-sumarProductos(productos...)
	b := <-sumarServicios(servicios...)
	c := <-sumarMantenimiento(mantenimiento...)

	fmt.Printf("Resultado: %.2f\n", sumar(a, b, c))

}

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int64
}

type Servicio struct {
	Nombre  string
	Precio  float64
	MinTrab float64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumar(a, b, c float64) float64 {
	fmt.Printf("a = %2.f\n, b = %2.f\n, c = %2.f\n", a, b, c)
	return a + b + c
}

func sumarProductos(productos ...Productos) <-chan float64 {

	c := make(chan float64)

	go func() {

		var suma float64

		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Productos ha comenzado")

		for _, producto := range productos {
			suma = suma + producto.Precio
		}
		c <- suma

	}()

	return c

}

func sumarServicios(servicios ...Servicio) <-chan float64 {

	c := make(chan float64)

	go func() {

		var suma float64

		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Servicios ha comenzado")

		for _, servicio := range servicios {
			suma = suma + servicio.Precio*(servicio.MinTrab/30)
		}

		c <- suma

	}()

	return c
}

func sumarMantenimiento(mantenimiento ...Mantenimiento) <-chan float64 {

	c := make(chan float64)
	var suma float64

	go func() {

		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Mantenimiento ha comenzado")

		for _, man := range mantenimiento {
			suma = suma + man.Precio
		}

		c <- suma
	}()

	return c
}
