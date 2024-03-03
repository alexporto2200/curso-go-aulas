package main

import "fmt"

func showType(i interface{}) {
	fmt.Printf("%v é do tipo %T\n", i, i)
}

func main() {

	var alex interface{} = "Alex"

	println(alex)

	println(alex.(string))

}
