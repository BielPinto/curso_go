package tax

func CalcalteTax(amount float64) float64 {
	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}
