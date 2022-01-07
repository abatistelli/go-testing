package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	a := 3
	b := 2
	resultadoEsperado := 1

	resultado := Restar(a, b)

	assert.Equal(t, resultadoEsperado, resultado, "Los Resultados deben ser iguales")
}
