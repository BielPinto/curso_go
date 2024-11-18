package main

type MyNumber int
type Number interface {
	~int | float64
}

// func Soma[T int | float64](m map[string]T) T {
func Sum[T Number](m map[string]T) T {

	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Gabriel": 150, "kal": 300, "Nina": 7}
	m2 := map[string]float64{"Gabriel": 157.8, "kal": 300.5, "Nina": 7.2}
	m3 := map[string]MyNumber{"Gabriel": 150, "kal": 300, "Nina": 7}

	println(Sum(m))
	println(Sum(m2))
	println(Sum(m3))
	println(compare(10, 20))
}
