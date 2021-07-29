package main

import "fmt"

type Polynomial struct {
	// Coefficients : c[0] + c[1]x + c[2]x^2 + c[3]x^3 + ...
	c []float64
}

func (p Polynomial) Plus(x Polynomial) (r Polynomial) {
	if len(p.c) > len(x.c) {
		r.c = p.c // Call by value = might be sent by copying "p.c" to "r.c"
		for i := range x.c {
			r.c[i] = r.c[i] + x.c[i]
		}
	} else {
		r.c = x.c // Call by value = might be sent by copying "p.c" to "r.c"
		for i := range p.c {
			r.c[i] = r.c[i] + p.c[i]
		}
	}
	return
}

func main() {
	p0 := Polynomial{c: []float64{1, 2, 3, 4, 5}}
	p_big := Polynomial{c: []float64{8, 7, 9, 4, 3, 1}}
	p_small := Polynomial{c: []float64{5, 7, 2, 3}}
	calcBig := p0.Plus(p_big)
	fmt.Println("  calcBig      (p0.c): ", p0.c) // No change of p0.c
	fmt.Println("  calcBig   (calcBig): ", calcBig.c)
	calcSmall := p0.Plus(p_small)
	fmt.Println("calcSmall      (p0.c): ", p0.c) // Why change???
	fmt.Println("calcSmall (calcSmall): ", calcSmall.c)
}
