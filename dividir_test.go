package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	a := 3
	b := 0

	_, err := Dividir(a, b)

	assert.Nil(t, err)

}
