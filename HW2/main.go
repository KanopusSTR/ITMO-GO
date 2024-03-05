package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func getCharByIndex(str string, idx int) rune {
	return []rune(str)[idx]
}

func getStringBySliceOfIndexes(str string, indexes []int) string {
	var st []rune

	var runes = []rune(str)
	for _, i := range indexes {
		st = append(st, runes[i])
	}
	return string(st)
}

func addPointers(ptr1, ptr2 *int) *int {
	if ptr1 != nil && ptr2 != nil {
		var sum = *ptr1 + *ptr2
		return &sum
	} else {
		return nil
	}
}

func isComplexEqual(a, b complex128) bool {
	return math.Abs(real(a)-real(b)) < 0.0001 && math.Abs(imag(a)-imag(b)) < 0.0001
}

func getRootsOfQuadraticEquation(a, b, c float64) (complex128, complex128) {
	var d = b*b - 4*a*c
	compB := complex(b, 0)
	var sqrtD = cmplx.Sqrt(complex(d, 0))
	return (sqrtD - compB) / 2, (-sqrtD - compB) / 2
}

func mergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}
	s1 := mergeSort(s[:len(s)/2])
	s2 := mergeSort(s[len(s)/2:])

	var ans []int

	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			ans = append(ans, s1[i])
			i++
		} else {
			ans = append(ans, s2[j])
			j++
		}
	}
	for ; i < len(s1); i++ {
		ans = append(ans, s1[i])
	}
	for ; j < len(s2); j++ {
		ans = append(ans, s2[j])
	}
	return ans
}

func reverseSliceOne(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func reverseSliceTwo(s []int) []int {
	ans := []int{}
	for i := len(s) - 1; i > -1; i-- {
		ans = append(ans, s[i])
	}
	return ans
}

func swapPointers(a, b *int) {
	*a, *b = *b, *a
}

func swapPointers2(a, b **int) {
	*a, *b = *b, *a
}

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func deleteByIndex(s []int, idx int) []int {
	slice1 := s[:idx]
	return append(slice1, s[idx+1:]...)
}

func main() {
	i, j := 1, 2
	a, b := &i, &j
	swapPointers2(&a, &b)
	fmt.Println(i, j, *a, *b)

}
