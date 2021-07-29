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
	if n <= 10 {
		if normalization {
			polynomials = []Polynomial{
				{[]float64{0.7071067811865476}},
				{[]float64{0, 1.2247448713915892}},
				{[]float64{-0.7905694150420949, 0, 2.371708245126285}},
				{[]float64{0, -2.8062430400804566, 0, 4.677071733467428}},
				{[]float64{0.795495128834866, 0, -7.954951288348662, 0, 9.280776503073438}},
				{[]float64{0, 4.397264774834466, 0, -20.52056894922751, 0, 18.46851205430476}},
				{[]float64{-0.7967217989988726, 0, 16.73115777897633, 0, -50.19347333692898, 0, 36.80854711374793}},
				{[]float64{0, -5.990715472712755, 0, 53.9164392544148, 0, -118.61616635971255, 0, 73.42905536553636}},
				{[]float64{0.7972004543733809, 0, -28.699216357441717, 0, 157.84568996592944, 0, -273.59919594094436, 0, 146.57099782550594}},
				{[]float64{0, 7.585118792715734, 0, -111.24840895983077, 0, 433.86879494334, 0, -619.8125642047713, 0, 292.689266430031}},
				{[]float64{-0.7974348906244046, 0, 43.85891898434226, 0, -380.1106311976329, 0, 1140.3318935928987, 0, -1384.688727934234, 0, 584.6463517944546}},
			}
			polynomials = polynomials[:(n + 1)]
			return
		} else {
			polynomials = []Polynomial{
				{[]float64{1}},
				{[]float64{0, 1}},
				{[]float64{-0.5, 0, 1.5}},
				{[]float64{0, -1.5, 0, 2.5}},
				{[]float64{0.375, 0, -3.75, 0, 4.375}},
				{[]float64{0, 1.875, 0, -8.75, 0, 7.875}},
				{[]float64{-0.3125, 0, 6.5625, 0, -19.6875, 0, 14.4375}},
				{[]float64{0, -2.1875, 0, 19.6875, 0, -43.3125, 0, 26.8125}},
				{[]float64{0.273438, 0, -9.84375, 0, 54.140625, 0, -93.84375, 0, 50.273438}},
				{[]float64{0, 2.460938, 0, -36.09375, 0, 140.765625, 0, -201.09375, 0, 94.9609388}},
				{[]float64{-0.246094, 0, 13.535156, 0, -117.304688, 0, 351.914062, 0, -427.324219, 0, 180.425781}},
			}
			polynomials = polynomials[:(n + 1)]
			return
		}
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
