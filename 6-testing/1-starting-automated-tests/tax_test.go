package tax

import "testing"

func TestCalculateTax(t *testing.T) {

	amount := 600.0
	expected := 6.0

	result := CalcalteTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
