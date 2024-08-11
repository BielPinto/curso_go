package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {

	amount := 600.0
	expected := 5.0

	result := CalcalteTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {

	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalcalteTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}

}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalcalteTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalcalteTax2(500.0)
	}
}
