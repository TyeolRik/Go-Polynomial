package polynomial

import (
	"fmt"
	"math"
	"sort"
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
	x := []float64{0.1819336, 0.4222107, 0.7173508, 0.06806584, 0.336313, 0.5157331, 0.6954517, 0.7335832, 0.3385519, 0.2617832}
	sort.Float64s(x)
	pos := LegendrePolynomials(5, true)
	k := 5
	for i := 0; i < k; i++ {
		tempSum := 0.0
		for _, eachX := range x {
			tempSum += pos[i+1].Evaluate(2*eachX-1) * math.Sqrt2
		}
		tempSum = tempSum / math.Sqrt(float64(len(x)))
		tempSum = tempSum * tempSum
		fmt.Println(i+1, tempSum)
	}
}
