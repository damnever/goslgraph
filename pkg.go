// Functions copy from https://github.com/cpmech/gosl
// Lisenced under BSD 3-Clause.

package goslgraph

import (
	"fmt"
)

var (
	verbose = false
)

// "github.com/cpmech/gosl/io"

// Sf wraps Sprintf
func Sf(msg string, prm ...interface{}) string {
	return fmt.Sprintf(msg, prm...)
}

// Pf prints formatted string
func Pf(msg string, prm ...interface{}) {
	if verbose {
		fmt.Printf(msg, prm...)
	}
}

// PfYel prints formatted string in high intensity yello
func PfYel(msg string, prm ...interface{}) {
	if verbose {
		fmt.Printf(msg, prm...)
	}
}

func Pforan(msg string, prm ...interface{}) {
	if verbose {
		fmt.Printf(msg, prm...)
	}
}

// "github.com/cpmech/gosl/chk"

// PrintTitle returns the Test Title
func PrintTitle(title string) {
	fmt.Printf("   . . . testing . . .   %s\n", title)
}

// "github.com/cpmech/gosl/utl"

// Imax returns the maximum between two integers
func Imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the minimum between two float point numbers
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// Alloc allocates a slice of slices of float64
func Alloc(m, n int) (mat [][]float64) {
	mat = make([][]float64, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]float64, n)
	}
	return
}

// IntAlloc allocates a matrix of integers
func IntAlloc(m, n int) (mat [][]int) {
	mat = make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}
	return
}
