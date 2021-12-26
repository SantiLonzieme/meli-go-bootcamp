package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("")

	ejercicioA()
	ejercicioB(-8, -5, 10, 5)
	ejercicioC(3000, "C")
	ejercicioD()
	ejercicioE()

}

func ejercicioD() {

	maxFunc, err := tipoDeOperacion(minimo)

	if err != nil {
		fmt.Println("El tipo de operación no existe")
		return
	}

	valorMax := maxFunc(5, 3, 6, 7, 10)

	fmt.Println(valorMax, err)
}

func ejercicioA() (impuesto int) {

	sueldo := 60000

	if sueldo > 50000 && sueldo < 150000 {

		impuesto := sueldo / 100 * 17

		fmt.Println("El empleado pagara", impuesto, "de impuestos")

		return impuesto

	} else {

		impuesto := sueldo / 100 * (17 + 10)

		fmt.Println("El empleado pagara", impuesto, "de impuestos")

		return impuesto

	}

}

func ejercicioB(notas ...int) (int, error) {

	var notasTotal int
	var promedio int

	for _, nota := range notas {

		if nota < 0 {

			fmt.Println("Hay un numero negativo")

			return 0, errors.New("Hay un numero negativo")
		}

		notasTotal = notasTotal + nota

		promedio = notasTotal / len(notas)

	}
	fmt.Println("El promedio del alumno es", promedio)

	return promedio, nil
}

func ejercicioC(min int, categoria string) (salario int) {

	switch categoria {

	case "C":

		salario = min / 60 * 1000

		fmt.Println(salario)

		return salario

	case "B":

		salario = min / 60 * 1500
		salario := salario + (salario / 100 * 20)

		fmt.Println(salario)

		return salario

	case "A":

		salario = min / 60 * 3000
		salario := salario + (salario / 100 * 50)

		fmt.Println(salario)

		return salario
	}
	return
}

const (
	minimo   = "minimo"
	maximo   = "maximo"
	promedio = "promedio"
)

func min(calificaciones ...int) (calificacion int) {

	calificacion = 10

	for _, nota := range calificaciones {

		if nota < calificacion {
			calificacion = nota
		}
	}

	fmt.Println("La nota más baja es", calificacion)

	return
}

func max(calificaciones ...int) (calificacion int) {

	calificacion = 0

	for _, nota := range calificaciones {

		if nota > calificacion {
			calificacion = nota
		}
	}

	fmt.Println("La nota más alta es", calificacion)

	return
}

func prom(calificaciones ...int) (calificacion int) {

	for _, nota := range calificaciones {

		calificacion = calificacion + nota
	}

	fmt.Println("El promedio es", calificacion)

	return
}

func tipoDeOperacion(operador string) (func(calificaciones ...int) (calificacion int), error) {

	switch operador {

	case "minimo":
		return min, nil
	case "maximo":
		return max, nil
	case "promedio":
		return prom, nil
	default:
		return nil, errors.New("No existe ese calculo")
	}
}

func ejercicioE() {

	funcAnimal, err := tipoDeAlimento("elefante")

	if err != nil {
		fmt.Println("El animal no existe")
		return
	}

	cantidadDeAlimento := funcAnimal(6)

	fmt.Println(cantidadDeAlimento, err)

}

func tipoDeAlimento(operador string) (func(animales int) (alimentoKg int), error) {

	switch operador {
	case "perro":
		return perroFunc, nil
	case "gato":
		return gatoFunc, nil
	case "hamster":
		return hamsterFunc, nil
	case "tarantula":
		return tarantulaFunc, nil
	default:
		return nil, errors.New("\nNo existe el animal")

	}
}

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func perroFunc(animales int) (alimentoKg int) {

	alimentoKg = animales * 10

	fmt.Println("Se necesitan para el perro", alimentoKg, "kg")

	return
}

func gatoFunc(animales int) (alimentoKg int) {

	alimentoKg = animales * 5

	fmt.Println("Se necesitan para el gato", alimentoKg, "kg")

	return
}

func hamsterFunc(animales int) (alimentoKg int) {

	alimentoKg = animales * 250

	fmt.Println("Se necesitan para el hamster", alimentoKg, "gramos")

	return
}

func tarantulaFunc(animales int) (alimentoKg int) {

	alimentoKg = animales * 150

	fmt.Println("Se necesitan para la tarantula", alimentoKg, "gramos")

	return
}
