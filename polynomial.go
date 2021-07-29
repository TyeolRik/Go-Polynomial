package polynomial

import (
	"log"
)

type Polynomial struct {
	c []float64 // Coefficients : a + bx + cx^2 + dx^3 -> {a, b, c, d}
}

// Create New object of Polynomial
func NewPolynomial(c []float64) (p Polynomial) {
	p.c = c
	return
}

func (p Polynomial) Coefficients() []float64 {
	return p.c
}

func (p Polynomial) Plus(x float64) (r Polynomial) {
	r.c = make([]float64, len(p.c))
	copy(r.c, p.c)
	for i := range r.c {
		r.c[i] = r.c[i] + x
	}
	return
}

func (p Polynomial) Minus(x float64) (r Polynomial) {
	r.c = make([]float64, len(p.c))
	copy(r.c, p.c)
	for i := range r.c {
		r.c[i] = r.c[i] - x
	}
	return
}

func (p Polynomial) Muliply(x float64) (r Polynomial) {
	r.c = make([]float64, len(p.c))
	copy(r.c, p.c)
	for i := range r.c {
		r.c[i] = r.c[i] * x
	}
	return
}

func (p Polynomial) Divide(x float64) (r Polynomial) {
	r.c = make([]float64, len(p.c))
	copy(r.c, p.c)
	for i := range r.c {
		r.c[i] = r.c[i] / x
	}
	return
}

func (p Polynomial) PlusPolynomial(x Polynomial) (r Polynomial) {
	if len(p.c) > len(x.c) {
		r.c = make([]float64, len(p.c))
		copy(r.c, p.c)
		for i := range x.c {
			r.c[i] = r.c[i] + x.c[i]
		}
	} else {
		r.c = make([]float64, len(x.c))
		copy(r.c, x.c)
		for i := range p.c {
			r.c[i] = r.c[i] + p.c[i]
		}
	}
	return
}

func (p Polynomial) MinusPolynomial(x Polynomial) (r Polynomial) {
	if len(p.c) > len(x.c) {
		r.c = make([]float64, len(p.c))
		copy(r.c, p.c)
		for i := range x.c {
			r.c[i] = r.c[i] - x.c[i]
		}
	} else {
		r.c = make([]float64, len(x.c))
		copy(r.c, x.c)
		for i := range p.c {
			r.c[i] = r.c[i] - p.c[i]
		}
	}
	return
}

func (p Polynomial) MultiplyPolynomial(x Polynomial) (r Polynomial) {
	r.c = make([]float64, len(p.c)+len(x.c)-1)
	for i := 0; i < len(p.c); i++ {
		for j := 0; j < len(x.c); j++ {
			r.c[i+j] += p.c[i] * x.c[j]
		}
	}
	return
}

func (p Polynomial) Evaluate(x float64) (calculated float64) {
	power := 1.0
	calculated = 0.0
	for _, coefficient := range p.c {
		calculated += coefficient * power
		power = power * x
	}
	return
}

func (p Polynomial) Derivative(times ...int) (d_poly Polynomial) {
	switch len(times) {
	case 0:
		d_poly.c = make([]float64, len(p.c)-1)
		for i := 0; i < len(p.c)-1; i++ {
			d_poly.c[i] = p.c[i+1] * float64(i+1)
		}
		return
	case 1:
		d_poly = p
		for i := 0; i < times[0]; i++ {
			d_poly = d_poly.Derivative(1)
		}
		return
	default:
		log.Panic("Something is wrong. Derivative Parameter should be 0 or 1, but we got ", len(times), times)
		return
	}
}
