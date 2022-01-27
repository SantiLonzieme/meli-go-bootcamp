package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciZero(t *testing.T) {

	resultadoEsperado := 0
	resultado := Fibonacci(0)

	assert.Equal(t, resultadoEsperado, resultado)

}

func TestFibonacciResul(t *testing.T) {

	resultadoEsperado := 55
	resultado := Fibonacci(10)

	assert.Equal(t, resultadoEsperado, resultado)

}
