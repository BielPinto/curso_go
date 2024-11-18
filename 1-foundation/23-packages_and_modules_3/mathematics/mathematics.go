package mathematics

func Sum[T int | float64](a, b T) T {
	return a + b
}

type Car struct {
	Brand string
	test  string
}

func (c Car) Walk() string {
	return "Car to walking"
}
