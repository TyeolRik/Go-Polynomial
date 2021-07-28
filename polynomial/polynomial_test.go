package polynomial

import (
	"fmt"
	"testing"
)

func TestPolynomial(t *testing.T) {
	po := NewPolynomial([]float64{0, 1.732051})
	new := po.MultiplyPolynomial(NewPolynomial([]float64{1, 2, 3, 4, 5})).Divide(3)
	fmt.Println(new)
}

func TestFunc(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4}
	fmt.Println(a)
	a = a[:3]
	fmt.Println(a)
}

func TestLegendrePolynomials(t *testing.T) {
	pos := LegendrePolynomials(5, false)
	for _, po := range pos {
		fmt.Printf("%.6f\n", po.c)
	}
}

func TestPredict(t *testing.T) {
	x := -0.8638683
	pos := LegendrePolynomials(5, true)[5]
	fmt.Println(pos.Evaluate(x))
}
