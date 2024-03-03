package main

func main() {

	a := 10

	// usar & para pegar o endereco de memoria
	var ponteiro *int = &a

	// variavel, ponteiro, endereco de memoria
	println(a, ponteiro, *ponteiro)

	*ponteiro = 20

	println(a, ponteiro, *ponteiro)

	b := &a

	println(b, *b)
}
