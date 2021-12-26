package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	ejercicio2()
}

func ejercicio2() {

	doc := []byte("ID, PRECIO, CANTIDAD; 1, 1000, 5; 2, 2000, 2; 3, 1000, 5")

	os.WriteFile("./archivo2.csv", doc, 0644)

	archivo, err := os.ReadFile("./archivo2.csv")
	impr := strings.Split(string(archivo), ";")

	if err != nil {
		fmt.Println("Error en la lectura del archivo")
	}

	var total int

	for _, value := range impr {

		line := strings.Split(string(value), ",")

		var totalPrecioStr string
		var totalCantidadStr string
		var totalCantidadInt int
		var totalPrecioInt int
		var errPrecio error
		var errCantidad error

		if line[1] != " PRECIO" || line[2] != " CANTIDAD" {
			totalPrecioStr = strings.Replace(line[1], " ", "", -1)
			totalCantidadStr = strings.Replace(line[2], " ", "", -1)

			totalPrecioInt, errPrecio = strconv.Atoi(totalCantidadStr)
			totalCantidadInt, errCantidad = strconv.Atoi(totalPrecioStr)

			total = total + totalPrecioInt*totalCantidadInt
		}

		if errPrecio != nil || errCantidad != nil {
			fmt.Println(errPrecio, errCantidad)
		}

		for i, value2 := range line {
			fmt.Printf("%s\t\t", value2)
			if i == len(line)-1 {
				fmt.Print("\n")
			}
		}
	}

	fmt.Printf("El total es %d\n", total)

}
