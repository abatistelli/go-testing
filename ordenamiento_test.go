package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	var slicesInteger = []int{6, 4, 1, 2, 9, 7, 5}

	resultadoEsperado := []int{1, 2, 4, 5, 6, 7, 9}

	resultado := ordenarSlice(slicesInteger)

	assert.Equal(t, resultadoEsperado, resultado, "Los output tienen que ser iguales")
}
