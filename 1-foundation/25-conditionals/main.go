package main

func main() {
	a := 1
	b := 2

	if a == 0 && b == 0 || b > a {
		println(a)
	}
	if a > +b {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("d")

	}
}
