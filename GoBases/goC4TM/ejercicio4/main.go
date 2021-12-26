package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(salarioMensual(200, 500))
	fmt.Println(medioAguinaldo(2000, 100))
	fmt.Println(salarioMensual(-200, 500))
	fmt.Println(medioAguinaldo(-2000, 100))

}

func salarioMensual(horasTrabajadas, valorHora float64) (float64, error) {

	var salario float64

	salario = horasTrabajadas * valorHora

	if salario >= 150000 {
		salario = salario - (salario / 100 * 10)
	}

	if horasTrabajadas < 80 || horasTrabajadas < 0 {

		fmt.Println(errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales"))

	}

	return salario, nil
}

func medioAguinaldo(mejorSalario, mesesTrabajados float64) (float64, error) {

	var aguinaldo float64

	aguinaldo = mejorSalario / 12 * mesesTrabajados

	if mejorSalario < 0 || mesesTrabajados < 0 {
		fmt.Println(errors.New("Hay un numero negativo"))
	}

	return aguinaldo, nil

}
