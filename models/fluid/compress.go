package main

import (
	"fmt"
	"math"
)

// TODO: to complete functions

func SettlePress(press, fractions []float64) float64 {
	// TODO: Description
	if len(press) != len(fractions) {
		// TODO: return error
		break

	} else {
		sum := 0.0
		for idx, _ := range press {
			sum = sum + press[idx]*fractions[idx]
		}
		return sum
	}
}

func VelSound(k, MW, TK float64) float64 {
	const R = 8314.0
	return math.Pow(k*R*TK/MW, 0.5)
}

func FlowOrifice(P, k, Mw, T, Vol float64) float64 {
	// All unit must be absolute
	const R = 8314.0
	a := (k + 1) / (k - 1)
	b := 2.0 / (k + 1)
	c1 := k * Mw / (R * T)
	c2 := math.Pow(b, a) * c1
	return (P) * 98066.5 * math.Pow(c2, 0.5)
}