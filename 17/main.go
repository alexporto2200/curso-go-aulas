package main

func Soma(n map[string]int) int {
	var soma int
	for _, v := range n {
		soma += v
	}

	return soma
}

func SomaGenerica[T int | float64](n map[string]T) T {
	var soma T
	for _, v := range n {
		soma += v
	}

	return soma
}

func main() {

	a := map[string]int{
		"Alex":  10,
		"João":  20,
		"Maria": 30,
	}

	b := map[string]float64{
		"Alex":  10.5,
		"João":  20.5,
		"Maria": 30.5,
	}
	println(Soma(a))
	println(SomaGenerica(a))

	println(SomaGenerica(b))
}
