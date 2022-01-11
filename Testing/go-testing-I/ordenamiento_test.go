package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {

	resultadoEsperado := []int{1, 2, 3, 4, 5, 6}

	resultado := Ordenamiento(2, 3, 1, 4, 6, 5)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
