// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goslgraph

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_munkres01(t *testing.T) {
	PrintTitle("munkres01")

	C := [][]float64{
		{1, 2, 3},
		{2, 4, 6},
		{3, 6, 9},
	}
	Ccor := [][]float64{
		{0, 1, 2},
		{0, 2, 4},
		{0, 3, 6},
	}
	Mcor := [][]MaskType{
		{StarType, NoneType, NoneType},
		{NoneType, NoneType, NoneType},
		{NoneType, NoneType, NoneType},
	}

	var mnk Munkres
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)

	// 1:
	PfYel("1: after step 0:\n")
	Pf("%v", mnk.StrCostMatrix())

	// 2: step 1
	nextStep := mnk.step1()
	PfYel("\n2: after step 1:\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 2, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{false, false, false}, mnk.colCovered, "col_covered")

	// 3: step 2
	nextStep = mnk.step2()
	PfYel("\n3: after step 2:\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 3, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{false, false, false}, mnk.colCovered, "col_covered")

	// 4: step 3
	nextStep = mnk.step3()
	PfYel("\n4: after step 3:\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 4, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, false, false}, mnk.colCovered, "col_covered")

	// 5: step 4
	nextStep = mnk.step4()
	PfYel("\n5: after step 4:\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 6, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, false, false}, mnk.colCovered, "col_covered")

	// 6: step 6
	Ccor = [][]float64{
		{0, 0, 1},
		{0, 1, 3},
		{0, 2, 5},
	}
	nextStep = mnk.step6()
	PfYel("\n6: after step 6:\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 4, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, false, false}, mnk.colCovered, "col_covered")

	// 7: step 4 again (1)
	Mcor = [][]MaskType{
		{StarType, PrimeType, NoneType},
		{PrimeType, NoneType, NoneType},
		{NoneType, NoneType, NoneType},
	}
	nextStep = mnk.step4()
	PfYel("\n7: after step 4 again (1):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 5, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{true, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{false, false, false}, mnk.colCovered, "col_covered")

	// 8: step 5
	Mcor = [][]MaskType{
		{NoneType, StarType, NoneType},
		{StarType, NoneType, NoneType},
		{NoneType, NoneType, NoneType},
	}
	nextStep = mnk.step5()
	PfYel("\n8: after step 5:\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 3, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{false, false, false}, mnk.colCovered, "col_covered")

	// 9: step 3 again (1)
	nextStep = mnk.step3()
	PfYel("\n9: after step 3 again (1):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 4, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, true, false}, mnk.colCovered, "col_covered")

	// 10: step 4 again (2)
	nextStep = mnk.step4()
	PfYel("\n10: after step 4 again (2):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 6, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, true, false}, mnk.colCovered, "col_covered")

	// 11: step 6 again (1)
	Ccor = [][]float64{
		{0, 0, 0},
		{0, 1, 2},
		{0, 2, 4},
	}
	nextStep = mnk.step6()
	PfYel("\n11: after step 6 again (1):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 4, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, true, false}, mnk.colCovered, "col_covered")

	// 12: step 4 again (3)
	Mcor = [][]MaskType{
		{NoneType, StarType, PrimeType},
		{StarType, NoneType, NoneType},
		{NoneType, NoneType, NoneType},
	}
	nextStep = mnk.step4()
	PfYel("\n12: after step 4 again (3):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 6, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{true, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, false, false}, mnk.colCovered, "col_covered")

	// 13: step 6 again (2)
	Ccor = [][]float64{
		{1, 0, 0},
		{0, 0, 1},
		{0, 1, 3},
	}
	nextStep = mnk.step6()
	PfYel("\n13: after step 6 again (2):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 4, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{true, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, false, false}, mnk.colCovered, "col_covered")

	// 14: step 4 again (4)
	Mcor = [][]MaskType{
		{NoneType, StarType, PrimeType},
		{StarType, PrimeType, NoneType},
		{PrimeType, NoneType, NoneType},
	}
	nextStep = mnk.step4()
	PfYel("\n14: after step 4 again (4):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 5, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{true, true, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{false, false, false}, mnk.colCovered, "col_covered")

	// 15: step 5 again (1)
	Mcor = [][]MaskType{
		{NoneType, NoneType, StarType},
		{NoneType, StarType, NoneType},
		{StarType, NoneType, NoneType},
	}
	nextStep = mnk.step5()
	PfYel("\n15: after step 5 again (1):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 3, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{false, false, false}, mnk.colCovered, "col_covered")

	// 15: step 3 again (2)
	nextStep = mnk.step3()
	PfYel("\n15: after step 3 again (2):\n")
	Pf("%v", mnk.StrCostMatrix())
	require.Equal(t, 7, nextStep)
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	checkMaskMatrix(t, "M", mnk.M, Mcor)
	require.Equal(t, []bool{false, false, false}, mnk.rowCovered, "row_covered")
	require.Equal(t, []bool{true, true, true}, mnk.colCovered, "col_covered")
}

func Test_munkres02(t *testing.T) {
	PrintTitle("munkres02")

	C := [][]float64{
		{2, 1},
		{1, 1},
	}
	var mnk Munkres
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{1, 0}, "links") // 0 goes with 1 and 1 goes with 0
	assertFloat64(t, mnk.Cost, 2, "cost")

	C = [][]float64{
		{2, 2},
		{4, 3},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{0, 1}, "links") // 0 goes 0 and 1 goes with 1
	assertFloat64(t, mnk.Cost, 5, "cost")

	C = [][]float64{
		{2, 2},
		{1, 3},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{1, 0}, "links") // 0 goes 1 and 1 goes 0
	assertFloat64(t, mnk.Cost, 3, "cost")

	C = [][]float64{
		{2, 1},
		{2, 1},
		{1, 1},
		{1, 1},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{-1, -1, 1, 0}, "links")
	assertFloat64(t, mnk.Cost, 2, "cost")

	C = [][]float64{
		{1, 2, 3},
		{6, 5, 4},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{0, 2}, "links") // 0 goes with 0 and 1 goes with 2
	assertFloat64(t, mnk.Cost, 5, "cost")

	C = [][]float64{
		{1, 2, 3},
		{6, 5, 4},
		{1, 1, 1},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{0, 2, 1}, "links") // 0 goes with 0, 1 goes with 2 and 2 goes with 1
	assertFloat64(t, mnk.Cost, 6, "cost")

	C = [][]float64{
		{2, 4, 7, 9},
		{3, 9, 5, 1},
		{8, 2, 9, 7},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{0, 3, 1}, "links")
	assertFloat64(t, mnk.Cost, 5, "cost")

	C = [][]float64{
		{1, 2, 3},
		{2, 4, 6},
		{3, 6, 9},
	}
	Ccor := [][]float64{
		{1, 0, 0},
		{0, 0, 1},
		{0, 1, 3},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	deep2(t, "C", 1e-17, mnk.C, Ccor)
	require.Equal(t, mnk.Links, []int{2, 1, 0}, "links") // 0 goes with 2, 1 goes with 1 and 2 goes with 0
	assertFloat64(t, mnk.Cost, 10, "cost")

	// from https://projecteuler.net/index.php?section=problems&id=345
	C = [][]float64{
		{7, 53, 183, 439, 863},
		{497, 383, 563, 79, 973},
		{287, 63, 343, 169, 583},
		{627, 343, 773, 959, 943},
		{767, 473, 103, 699, 303},
	}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(C[i]); j++ {
			C[i][j] *= -1
		}
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	require.Equal(t, mnk.Links, []int{4, 1, 2, 3, 0}, "links")
}

func Test_munkres03(t *testing.T) {
	PrintTitle("munkres03. Euler problem 345")

	C := [][]float64{
		{7, 53, 183, 439, 863, 497, 383, 563, 79, 973, 287, 63, 343, 169, 583},
		{627, 343, 773, 959, 943, 767, 473, 103, 699, 303, 957, 703, 583, 639, 913},
		{447, 283, 463, 29, 23, 487, 463, 993, 119, 883, 327, 493, 423, 159, 743},
		{217, 623, 3, 399, 853, 407, 103, 983, 89, 463, 290, 516, 212, 462, 350},
		{960, 376, 682, 962, 300, 780, 486, 502, 912, 800, 250, 346, 172, 812, 350},
		{870, 456, 192, 162, 593, 473, 915, 45, 989, 873, 823, 965, 425, 329, 803},
		{973, 965, 905, 919, 133, 673, 665, 235, 509, 613, 673, 815, 165, 992, 326},
		{322, 148, 972, 962, 286, 255, 941, 541, 265, 323, 925, 281, 601, 95, 973},
		{445, 721, 11, 525, 473, 65, 511, 164, 138, 672, 18, 428, 154, 448, 848},
		{414, 456, 310, 312, 798, 104, 566, 520, 302, 248, 694, 976, 430, 392, 198},
		{184, 829, 373, 181, 631, 101, 969, 613, 840, 740, 778, 458, 284, 760, 390},
		{821, 461, 843, 513, 17, 901, 711, 993, 293, 157, 274, 94, 192, 156, 574},
		{34, 124, 4, 878, 450, 476, 712, 914, 838, 669, 875, 299, 823, 329, 699},
		{815, 559, 813, 459, 522, 788, 168, 586, 966, 232, 308, 833, 251, 631, 107},
		{813, 883, 451, 509, 615, 77, 281, 613, 459, 205, 380, 274, 302, 35, 805},
	}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(C[i]); j++ {
			C[i][j] *= -1
		}
	}

	var mnk Munkres
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	Pforan("links = %v\n", mnk.Links)
	Pforan("cost = %v  (13938)\n", -mnk.Cost)
	require.Equal(t, mnk.Links, []int{9, 10, 7, 4, 3, 0, 13, 2, 14, 11, 6, 5, 12, 8, 1}, "links")
	assertFloat64(t, mnk.Cost, -13938, "cost")
}

func checkMaskMatrix(t *testing.T, msg string, res, correct [][]MaskType) {
	if len(res) != len(correct) {
		Pf("%s. len(res)=%d != len(correct)=%d\n", msg, len(res), len(correct))
		t.Errorf("%s failed: res and correct matrices have different lengths. %d != %d", msg, len(res), len(correct))
		return
	}
	for i := 0; i < len(res); i++ {
		if len(res[i]) != len(correct[i]) {
			Pf("%s. len(res[%d])=%d != len(correct[%d])=%d\n", msg, i, len(res[i]), i, len(correct[i]))
			t.Errorf("%s failed: matrices have different number of columns", msg)
			return
		}
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] != correct[i][j] {
				Pf("[%d,%d] %v != %v\n", i, j, res[i][j], correct[i][j])
				t.Errorf("%s failed: different int matrices:\n [%d,%d] item is wrong: %v != %v", msg, i, j, res[i][j], correct[i][j])
				return
			}
		}
	}
}

func Test_munkres04(t *testing.T) {

	//verbose()
	PrintTitle("munkres04. row and column matrices")

	C := [][]float64{
		{1.0},
		{2.0},
		{0.5},
		{3.0},
		{4.0},
	}

	var mnk Munkres
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	Pforan("links = %v\n", mnk.Links)
	Pforan("cost = %v  (13938)\n", mnk.Cost)
	require.Equal(t, mnk.Links, []int{-1, -1, 0, -1, -1}, "links")
	assertFloat64(t, mnk.Cost, 0.5, "cost")

	C = [][]float64{
		{1.0, 2.0, 0.5, 3.0, 4.0},
	}
	mnk.Init(len(C), len(C[0]))
	mnk.SetCostMatrix(C)
	mnk.Run()
	Pforan("links = %v\n", mnk.Links)
	Pforan("cost = %v  (13938)\n", mnk.Cost)
	require.Equal(t, mnk.Links, []int{2}, "links")
	assertFloat64(t, mnk.Cost, 0.5, "cost")
}

func Test_munkres05(t *testing.T) {

	//verbose()
	PrintTitle("munkres05. issue (22/Nov/2016)") // fixed: use square matrix internally

	C := [][]float64{
		{11757.0, 6957.0},
		{28985.0, 24171.0},
		{33857.0, 29057.0},
	}

	D := [][]float64{
		{11757.0, 6957.0},
		{33857.0, 29057.0},
		{28985.0, 24171.0},
	}

	var mnkC Munkres
	mnkC.Init(len(C), len(C[0]))
	mnkC.SetCostMatrix(C)
	mnkC.Run()
	Pforan("C: links = %v\n", mnkC.Links)
	Pforan("C: cost = %v\n", mnkC.Cost)
	require.Equal(t, mnkC.Links, []int{0, 1, -1}, "C: links")
	assertFloat64(t, 35928, mnkC.Cost, "C: cost")

	Pf("\n")

	var mnkD Munkres
	mnkD.Init(len(D), len(D[0]))
	mnkD.SetCostMatrix(D)
	mnkD.Run()
	Pforan("D: links = %v\n", mnkD.Links)
	Pforan("D: cost = %v\n", mnkD.Cost)
	require.Equal(t, []int{0, -1, 1}, mnkD.Links, "D: links")
	assertFloat64(t, 35928, mnkD.Cost, "D: cost")
}

func assertFloat64(t *testing.T, a, b float64, msg ...interface{}) {
	require.True(t, math.Abs(a-b) < 1e-17)
}

func deep2(t *testing.T, msg string, tol float64, a, b [][]float64) {
	zero := false
	if len(b) == 0 {
		zero = true
	} else {
		if len(a) != len(b) {
			t.Fatalf("%s len(a)=%d != len(b)=%d", msg, len(a), len(b))
			return
		}
	}
	for i := 0; i < len(a); i++ {
		if !zero {
			if len(a[i]) != len(b[i]) {
				t.Fatalf("%s len(a[%d])=%d != len(b[%d])=%d", msg, i, len(a[i]), i, len(b[i]))
				return
			}
		}
		for j := 0; j < len(a[i]); j++ {
			var c float64
			if !zero {
				c = b[i][j]
			}
			if tstDiff(t, msg+fmt.Sprintf(" [%d][%d] ", i, j), tol, a[i][j], c, false) {
				return
			}
		}
	}
}

func tstDiff(t *testing.T, msg string, tol, a, b float64, showOK bool) (failed bool) {
	diff := math.Abs(a - b)
	if math.IsNaN(diff) || math.IsInf(diff, 0) {
		t.Fatalf("%s NaN or Inf in a=%v b=%v", msg, a, b)
		return true
	}
	if diff > tol {
		t.Fatalf("%s %v != %v |diff| = %g", msg, a, b, diff)
		return true
	}
	return
}
