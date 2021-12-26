package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Inicia")

	_, err := os.Open("./archivo.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println("Termino")
}
