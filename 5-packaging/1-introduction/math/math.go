package main

type Math struct {
	A int
	B int
}

func (m Math) ADD() int {
	return m.A + m.B

}
