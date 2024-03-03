package main

import "fmt"

func showType(i interface{}) {
	fmt.Printf("%v é do tipo %T\n", i, i)
}

func main() {

	//evitar utilizar interfaces vazias
	//elas aceitam qualquer tipo
	//e podem causar panics

	var x interface{} = 10
	x = "tttt"

	//println(x)

	var y interface{} = "Olá"

	//println(y)

	showType(x)
	showType(y)
}
