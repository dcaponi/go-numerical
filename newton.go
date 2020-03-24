package numerical

import (
	"errors"
	"fmt"
	"math"
)

// Solves nonlinear equations using Newton's Method

// Newton accepts an initial guess, an error threshold er and a function to solve.
func Newton(x float64, er float64, iter int, fxn func(x float64) float64, dfxn func(x float64) float64) (float64, error) {
	var s float64
	i := 0
	for i < iter {
		xn := x - (fxn(x) / dfxn(x))
		if (math.Abs(xn-x) / x) < er {
			s = xn
			fmt.Printf("finished after %d iterations\n", i)
			break
		}
		i++
		x = xn
	}
	if i >= iter {
		return s, errors.New("max iterations reached")
	}
	return s, nil
}
