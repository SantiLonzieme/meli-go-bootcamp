package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	fmt.Println("")

	cliente := NuevoCliente{}

	id := cliente.generarID(5)
	cliente.clienteExistente()

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		defer fmt.Println("No han quedado archivos abiertos")
		defer fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		defer fmt.Println("Fin de la ejecución")
	}()

	if id == 0 {
		panic("No se recibio ningun id")
	}

}

type NuevoCliente struct {
	Legajo          int
	NombreYApellido string
	DNI             int
	NumTelefono     int
	Domicilio       string
}

func (n *NuevoCliente) generarID(id int) int {

	if id != 0 {
		n.Legajo = rand.Int()
	}

	fmt.Println(n.Legajo)

	return id

}

func (n *NuevoCliente) clienteExistente() {

	_, err := os.Open("./documento.txt")

	defer func() {
		err := recover()

		fmt.Println("El archivo indicado no fue encontrado o está dañado\n", err)
	}()

	if err != nil {
		panic(err)
	}
}

func (n *NuevoCliente) verificacion(Nombre string, dni int, telefono int,
	domicilio string) (NuevoCliente, error) {

	cliente := NuevoCliente{n.Legajo, Nombre, dni, telefono, domicilio}
	var err error

	if cliente.NombreYApellido == "" || cliente.DNI == 0 || cliente.Domicilio == "" || cliente.NumTelefono == 0 {
		fmt.Println("Dato no ingresado")
	}

	return cliente, err
}
