package main

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestMyGetCharByIndex(t *testing.T) {
	str := "abcdefgh强幸i"
	require.Equal(t, 'a', getCharByIndex(str, 0))
	require.Equal(t, 'b', getCharByIndex(str, 1))
	require.Equal(t, 'c', getCharByIndex(str, 2))
	require.Equal(t, 'h', getCharByIndex(str, 7))
	require.Equal(t, '强', getCharByIndex(str, 8))
	require.Equal(t, '幸', getCharByIndex(str, 9))

}

func TestMyGetStringBySliceOfIndexes(t *testing.T) {
	str := "abcdefgh强幸i"
	require.Equal(t, "bcac", getStringBySliceOfIndexes(str, []int{1, 2, 0, 2}))
	require.Equal(t, "b强幸c", getStringBySliceOfIndexes(str, []int{1, 8, 9, 2}))
}

func TestMyAddPointers(t *testing.T) {
	num1 := 5
	num2 := 6
	num3 := 11
	require.Equal(t, &num3, addPointers(&num1, &num2))
}

func TestMyIsComplexEqual(t *testing.T) {
	var comp1 = complex(5, 1)
	var comp2 = complex(6, 1)
	var comp3 = complex(5, 2)
	var comp4 = complex(5, 1.0000000000000001)
	var comp5 = complex(5, 1)
	require.False(t, isComplexEqual(comp1, comp2))
	require.False(t, isComplexEqual(comp1, comp3))
	require.True(t, isComplexEqual(comp1, comp4))
	require.True(t, isComplexEqual(comp1, comp5))
}

func TestMyGetRootsOfQuadraticEquation(t *testing.T) {
	a, b := getRootsOfQuadraticEquation(1, 0, -2)
	require.Equal(t, complex(math.Sqrt(2), 0), a)
	require.Equal(t, complex(-math.Sqrt(2), 0), b)

	a, b = getRootsOfQuadraticEquation(1, 0, 2)
	require.Equal(t, complex(0, math.Sqrt(2)), a)
	require.Equal(t, complex(0, -math.Sqrt(2)), b)
}

func TestMyMergeSort(t *testing.T) {
	slice := []int{4, 2, 3, 1}
	sortedSlice := []int{1, 2, 3, 4}
	require.Equal(t, sortedSlice, mergeSort(slice))
}

func TestMyReverseSliceOne(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	reversedSlice := []int{4, 3, 2, 1}
	reverseSliceOne(slice)
	require.Equal(t, reversedSlice, slice)
}

func TestMyReverseSliceTwo(t *testing.T) {
	startSlice := []int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4}
	reversedSlice := []int{4, 3, 2, 1}
	require.Equal(t, reversedSlice, reverseSliceTwo(slice))
	require.Equal(t, startSlice, slice)
}

func TestMySwapPointers(t *testing.T) {
	num1 := 5
	num2 := 6
	swapPointers(&num1, &num2)
	require.Equal(t, 5, num2)
	require.Equal(t, 6, num1)
}

func TestMyIsSliceEqual(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{1, 2, 3, 4}
	slice3 := []int{4, 3, 2, 1}
	require.True(t, isSliceEqual(slice1, slice2))
	require.False(t, isSliceEqual(slice1, slice3))
}

func TestMyDeleteByIndex(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{1, 2, 4}
	slice3 := []int{1, 2}
	slice4 := []int{2}
	slice5 := []int{}
	require.Equal(t, slice2, deleteByIndex(slice1, 2))
	require.Equal(t, slice3, deleteByIndex(slice2, 2))
	require.Equal(t, slice4, deleteByIndex(slice3, 0))
	require.Equal(t, slice5, deleteByIndex(slice4, 0))
}
