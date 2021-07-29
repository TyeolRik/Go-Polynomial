package polynomial

import (
	"log"
	"math"
)

type dataFrame struct {
	c float64
	d float64
	e float64
	f float64
}

// legendre.recurrences in R.cran.orthopolynom
// GPL 3.0
// https://github.com/cran/orthopolynom/blob/master/R/legendre.recurrences.R
func recurrences(n int, normalize ...bool) (relations []dataFrame) {
	var normalization bool = false
	switch len(normalize) {
	case 0:
		normalization = false
	case 1:
		normalization = normalize[0]
	default:
		log.Panic("Something is wrong. the length of Normalize Parameter should be 0 or 1, but we got ", len(normalize), normalize)
	}

	if n < 0 {
		log.Panic("negative highest polynomial order")
	}

	relations = make([]dataFrame, (n + 1))
	j := 0

	switch normalization {
	case true:
		var j_float64 float64
		for j <= n {
			j_float64 = float64(j)
			var tempF float64
			if j_float64 == 0 {
				tempF = 0
			} else {
				tempF = j_float64 * math.Sqrt((2*j_float64+3)/(2*j_float64-1))
			}
			relations[j] = dataFrame{
				c: j_float64 + 1.0,
				d: 0.0,
				e: math.Sqrt((2*j_float64 + 1) * (2*j_float64 + 3)),
				f: tempF,
			}
			j = j + 1
		}
		return
	case false:
		var j_float64 float64
		for j <= n {
			j_float64 = float64(j)
			relations[j] = dataFrame{
				c: j_float64 + 1.0,
				d: 0.0,
				e: 2*j_float64 + 1,
				f: j_float64,
			}
			j = j + 1
		}
		return
	}
	return
}

// https://github.com/cran/orthopolynom/blob/master/R/orthonormal.polynomials.R
func orthonormalPolynomials(recurrences []dataFrame, p0 Polynomial) (polynomials []Polynomial) {
	np1 := len(recurrences)
	n := np1 - 1
	polynomials = make([]Polynomial, np1)
	polynomials[0] = p0
	j := 0
	for j < n {
		monomial := NewPolynomial([]float64{recurrences[j].d, recurrences[j].e})
		var p_jp1 Polynomial
		if j == 0 {
			p_jp1 = (monomial.MultiplyPolynomial(p0)).Divide(recurrences[j].c)
		} else {
			p_jp1 = (monomial.MultiplyPolynomial(polynomials[j]).MinusPolynomial(polynomials[j-1].Muliply(recurrences[j].f))).Divide(recurrences[j].c)
		}
		polynomials[j+1] = p_jp1
		j = j + 1
	}
	return
}

func orthogonalPolynomials(recurrences []dataFrame) (polynomials []Polynomial) {
	polynomials = orthonormalPolynomials(recurrences, NewPolynomial([]float64{1}))
	return
}

// https://github.com/cran/orthopolynom/blob/master/R/legendre.polynomials.R
func LegendrePolynomials(n int, normalize ...bool) (polynomials []Polynomial) {
	var normalization bool = false
	switch len(normalize) {
	case 0:
		normalization = false
	case 1:
		normalization = normalize[0]
	default:
		log.Panic("Something is wrong. the length of Normalize Parameter should be 0 or 1, but we got ", len(normalize), normalize)
	}

	if n < 0 {
		log.Panic("negative highest polynomial order")
	}
	recurrence := recurrences(n, normalization)
	switch normalization {
	case true:
		// h0 := 2.0
		p0 := Polynomial{c: []float64{1.0 / math.Sqrt2}}
		polynomials = orthonormalPolynomials(recurrence, p0)
		return
	case false:
		polynomials = orthogonalPolynomials(recurrence)
		return
	}
	return
}
