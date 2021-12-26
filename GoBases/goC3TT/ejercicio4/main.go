package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("")
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	ordInser(variable1)
	ordBurbuja(variable2)
	ordSeleccion(variable3)

	// ordeInsercion := <-ordInser(variable1)
	// ordeBurbuja := <-ordBurbuja(variable2)
	// ordeSeleccion := <-ordSeleccion(variable3)

	// fmt.Println("Burbuja\n", ordeBurbuja, "\n", "Inserción\n", ordeInsercion, "\n", "Seleccion\n", ordeSeleccion)

}

func ordInser(arr []int) <-chan []int {
	var auxiliar int
	channel := make(chan []int)
	start := time.Now()
	go func() {

		for i := 1; i < len(arr); i++ {
			auxiliar = arr[i]
			for j := i - 1; j >= 0 && arr[j] > auxiliar; j-- {
				arr[j+1] = arr[j]
				arr[j] = auxiliar
			}
		}

		channel <- arr
	}()
	tiempo := time.Since(start)
	fmt.Printf("Inserción tardó %s\n", tiempo)
	return channel

}

func ordSeleccion(arr []int) <-chan []int {

	channel := make(chan []int)
	start := time.Now()
	go func() {

		for i := 0; i < len(arr); i++ {
			minimo_encontrado, posicion_minimo := arr[i], i

			valor_original := arr[i]
			for j := i + 1; j < len(arr); j++ {
				valor_comparacion := arr[j]
				if valor_comparacion < minimo_encontrado {
					minimo_encontrado, posicion_minimo = valor_comparacion, j
				}
			}

			if minimo_encontrado != valor_original {
				arr[i], arr[posicion_minimo] = minimo_encontrado, valor_original
			}
		}

		channel <- arr

	}()
	tiempo := time.Since(start)
	fmt.Printf("Seleccion tardó %s\n", tiempo)

	return channel

}

func ordBurbuja(arr []int) <-chan []int {
	var auxiliar int
	channel := make(chan []int)
	start := time.Now()
	go func() {

		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if arr[i] > arr[j] {
					auxiliar = arr[i]
					arr[i] = arr[j]
					arr[j] = auxiliar
				}
			}
		}

		channel <- arr
	}()
	tiempo := time.Since(start)
	fmt.Printf("Burbuja tardó %s\n", tiempo)
	return channel

}
