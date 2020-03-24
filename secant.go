package numerical

import (
	"errors"
	"fmt"
	"math"
)

// Secant accepts 2 x points in the neighborhood of the possible solution, an error threshold er and a function to solve.
func Secant(x1 float64, x2 float64, er float64, iter int, fn func(x float64) float64) (float64, error) {
	i := 0
	var s float64

	for i < iter {
		f2 := fn(x2)
		n := x2 - ((f2 * (x1 - x2)) / (fn(x1) - f2))
		if (math.Abs(n-x2) / x2) < er {
			s = n
			fmt.Printf("complete after %d iterations", i)
			break
		}
		i++
		x1 = x2
		x2 = n
	}

	if i >= iter {
		return s, errors.New("max iterations reached")
	}
	return s, nil
}
