package main

import (
	"fmt"
)

func main() {

	salary := 130000

	if salary < 150000 {

		err := fmt.Errorf("error: el minimo no imponible es de 150000 y el salario ingresado es de %d", salary)
		fmt.Println(err)

		return
	}

	fmt.Println("Debe pagar impuesto")
}
