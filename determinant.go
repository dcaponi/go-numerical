package numerical

import (
	"errors"
)

// todo put this in a utility module
type stack []float64

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(n float64) {
	*s = append(*s, n)
}

func (s *stack) pop() (float64, bool) {
	if s.isEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	n := (*s)[i]
	*s = (*s)[:i]
	return n, true
}

func (s *stack) ToSlice() []float64 {
	return *s
}

// Det takes a 2D square matrix of floats and returns a float value for the determinant
func Det(mat [][]float64) (float64, error) {
	if len(mat) != len(mat[0]) {
		return 0.0, errors.New("determinant can only be performed on square matrices")
	}
	if len(mat) == 1 {
		return (mat[0][0]), nil
	}
	if len(mat) == 2 {
		return (mat[0][0] * mat[1][1]) - (mat[0][1] * mat[1][0]), nil
	}
	s := 0.0
	for i := 0; i < len(mat[0]); i++ {
		sm := subMat(mat[1:][:], i)
		z, err := Det(sm)
		if err == nil {
			if i%2 != 0 {
				s -= mat[0][i] * z
			} else {
				s += mat[0][i] * z
			}
		}
	}
	return s, nil
}

func subMat(mat [][]float64, p int) [][]float64 {
	stacks := make([]stack, len(mat))
	for n := range mat {
		stacks[n] = stack{}
		for j := range mat[n] {
			if j != p {
				stacks[n].push(mat[n][j])
			}
		}
	}
	out := make([][]float64, len(mat))
	for k := range stacks {
		out[k] = stacks[k].ToSlice()
	}
	return out
}
