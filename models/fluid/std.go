package std

import (
	"fmt"
	"math"
)

const (
	maxiter = 1e6
	tol     = 1e-8
)

func FactorFCole(re, eps, D float64) float64 {
	// TODO: add explanation

	if re <= 2300 {
		return 64.0 / re

	} else if re < 4000 {
		// i dont have a function for this Reynolds
		return 1000000000000000000000000000000
	} else {
		//	https://es.wikipedia.org/wiki/Ecuaci%C3%B3n_de_Colebrook-White
		iter := 1 // maximum iterations
		calc_error := 1.0
		ffactor := 0.001 //friction factor
		a := eps / D

		if calc_error > tol {
			for iter < int(maxiter) {
				x := 1.0 / math.Sqrt(ffactor)

				iter = iter + 1
			}
		}

	}
}
