package main

import "fmt"

func main() {

	palabra := "avion"

	fmt.Println("La palabra tiene", len(palabra), "letras")

	for i := 0; i < len(palabra); i++ {

		letra := palabra[i]

		fmt.Printf("%c \n", letra)
	}

	ejercicioB()
	ejercicioC()
	ejercicioD()
	ejercicioE()
	ejercicioF()
}

func ejercicioB() {

	precio := 50000.25
	porcentajeDescuento := 20.50

	descuento := precio / 100 * porcentajeDescuento

	fmt.Printf("El precio final es %.2f\n", precio-descuento)

}

func ejercicioC() {

	edad := 25
	empleado := true
	antiguedad := 2
	sueldo := 120.000

	if edad > 22 && empleado && antiguedad > 1 && sueldo > 100.000 {
		fmt.Println("El ciente puede acceder y no se le cobraran impuestos")
	} else if edad > 22 && empleado && antiguedad > 1 && sueldo < 100.000 {
		fmt.Println("El ciente puede acceder y se le cobraran impuestos")
	} else {
		fmt.Println("El ciente no puede acceder")
	}
}

func ejercicioD() {

	mes := 2

	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 7:
		fmt.Println("Junio")
	case 8:
		fmt.Println("Julio")
	case 9:
		fmt.Println("Agosto")
	case 10:
		fmt.Println("Septiembre")
	case 11:
		fmt.Println("Octubre")
	case 12:
		fmt.Println("Noviembre")
	case 13:
		fmt.Println("Diciembre")
	default:
		fmt.Println("El número de mes no existe")
	}
}

func ejercicioE() {

	var lista = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores",
		"Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	lista = append(lista, "Gabriela")

	fmt.Println(lista)
}

func ejercicioF() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	mayores := 0

	for _, edad := range employees {
		if edad >= 21 {
			mayores++
		}
	}

	fmt.Println(mayores, "son mayores de edad")

	employees["Federico"] = 25

	fmt.Println(employees, "Con federico")

	delete(employees, "Pedro")

	fmt.Println(employees, "Sin Pedro")

}
