package main

import (
	"fmt"
)

func main() {

	fmt.Println("")

	ejercicioA()
	ejercicioB(8, 5, 10, 5)
	ejercicioC(3000, "C")
	ejercicioD()
	ejercicioE()

}

func ejercicioD() {

	maxFunc := tipoDeOperacion("maximo")
	valorMax := maxFunc(5, 3, 6, 7, 10)

	fmt.Println(valorMax)
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

func ejercicioB(notas ...int) (promedio int) {

	var notasTotal int

	for _, nota := range notas {

		if nota < 0 {
			//errors.New("Hay un numero negativo")

			return
		}

		notasTotal = notasTotal + nota

		promedio = notasTotal / len(notas)

	}
	fmt.Println("El promedio del alumno es", promedio)

	return promedio
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

// func ejercicioE() {

// }

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

func tipoDeOperacion(operador string) func(calificaciones ...int) (calificacion int) {

	switch operador {

	case "minimo":
		return min
	case "maximo":
		return max
	case "promedio":
		return prom
	}

	return nil
}

func ejercicioE() {

	funcAnimal := tipoDeAlimento(perro)
	cantidadDeAlimento := funcAnimal(6)

	fmt.Println(cantidadDeAlimento)

}

func tipoDeAlimento(operador string) func(animales int) (alimentoKg int) {

	switch operador {
	case "perro":
		return perroFunc
	case "gato":
		return gatoFunc
	case "hamster":
		return hamsterFunc
	case "tarantula":
		return tarantulaFunc

	}

	return nil

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

	//pasarlo a kg

	fmt.Println("Se necesitan para el hamster", alimentoKg, "gramos")

	return
}

func tarantulaFunc(animales int) (alimentoKg int) {

	alimentoKg = animales * 150

	fmt.Println("Se necesitan para la tarantula", alimentoKg, "gramos")

	return
}
