package main

func main() {

	const a = 10
	const b = 20
	const c = 30
	const t = true
	const f = false

	if a > b {
		println("a > b")
	} else {
		println("a <= b")
	}

	if a > b && b < c {
		println("a > b e b < c")
	}

	if a > b || b < c {
		println("a > b ou b < c")
	}

	switch {
	case a > b:
		println("case a > b")
	case b < c:
		println("case b < c")
	default:
		println("case nenhum caso")
	}

	switch a {
	case 1:
		println("case 1")
	case 2:
		println("case 2")
	default:
		println("case nenhum caso")
	}

}
