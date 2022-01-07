package test

func ordenarSlice(sliceDesordenado []int) []int {
	var aux int
	for i := 0; i < len(sliceDesordenado); i++ {
		for j := 0; j < len(sliceDesordenado); j++ {
			if sliceDesordenado[i] < sliceDesordenado[j] {
				aux = sliceDesordenado[i]
				sliceDesordenado[i] = sliceDesordenado[j]
				sliceDesordenado[j] = aux
			}
		}
	}
	return sliceDesordenado
}
