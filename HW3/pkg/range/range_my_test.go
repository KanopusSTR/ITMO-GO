package rangeI

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLengthNormal(t *testing.T) {
	rangeInt := NewRangeInt(0, 5)
	require.Equal(t, 6, rangeInt.Length())
}

func TestMyLengthEmpty(t *testing.T) {
	rangeInt := NewRangeInt(1, 0)
	require.Equal(t, 0, rangeInt.Length())
}

func TestIntersectNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(0, 5)
	rangeInt2 := NewRangeInt(3, 10)
	rangeInt3 := NewRangeInt(3, 5)
	rangeInt1.Intersect(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestIntersectNo(t *testing.T) {
	rangeInt1 := NewRangeInt(0, 5)
	rangeInt2 := NewRangeInt(10, 12)
	rangeInt3 := NewRangeInt(1, 0)
	rangeInt1.Intersect(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestIntersectEmptyArg(t *testing.T) {
	rangeInt1 := NewRangeInt(0, 5)
	rangeInt2 := NewRangeInt(1, 0)
	rangeInt3 := NewRangeInt(1, 0)
	rangeInt1.Intersect(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestIntersectEmptyMain(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	rangeInt2 := NewRangeInt(1, 10)
	rangeInt3 := NewRangeInt(1, 0)
	rangeInt1.Intersect(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestUnionNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(0, 5)
	rangeInt2 := NewRangeInt(3, 10)
	rangeInt3 := NewRangeInt(0, 10)
	rangeInt1.Union(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestUnionNo(t *testing.T) {
	rangeInt1 := NewRangeInt(0, 5)
	rangeInt2 := NewRangeInt(10, 12)
	rangeInt3 := NewRangeInt(0, 5)
	rangeInt1.Union(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestUnionEmptyArg(t *testing.T) {
	rangeInt1 := NewRangeInt(0, 5)
	rangeInt2 := NewRangeInt(1, 0)
	rangeInt3 := NewRangeInt(0, 5)
	rangeInt1.Union(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestUnionEmptyMain(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	rangeInt2 := NewRangeInt(1, 10)
	rangeInt3 := NewRangeInt(1, 10)
	rangeInt1.Union(rangeInt2)
	require.Equal(t, rangeInt3, rangeInt1)
}

func TestIsEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	require.True(t, rangeInt1.IsEmpty())
	rangeInt2 := NewRangeInt(0, 1)
	require.False(t, rangeInt2.IsEmpty())
}

func TestContainsIntNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(6, 10)
	require.True(t, rangeInt1.ContainsInt(7))
	require.False(t, rangeInt1.ContainsInt(0))
}

func TestContainsIntEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	require.False(t, rangeInt1.ContainsInt(0))
	require.False(t, rangeInt1.ContainsInt(1))
}

func TestContainsRangeNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(6, 10)
	rangeInt2 := NewRangeInt(0, 5)
	rangeInt3 := NewRangeInt(7, 9)
	require.False(t, rangeInt1.ContainsRange(rangeInt2))
	require.True(t, rangeInt1.ContainsRange(rangeInt3))
}

func TestContainsRangeEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	rangeInt2 := NewRangeInt(0, 1)
	require.False(t, rangeInt1.ContainsRange(rangeInt2))
}

func TestIsIntersectNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(5, 10)
	rangeInt2 := NewRangeInt(3, 6)
	rangeInt3 := NewRangeInt(9, 12)
	rangeInt4 := NewRangeInt(0, 4)
	rangeInt5 := NewRangeInt(15, 20)
	require.True(t, rangeInt1.IsIntersect(rangeInt2))
	require.True(t, rangeInt1.IsIntersect(rangeInt3))
	require.False(t, rangeInt1.IsIntersect(rangeInt4))
	require.False(t, rangeInt1.IsIntersect(rangeInt5))
}

func TestIsIntersectEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	rangeInt2 := NewRangeInt(0, 0)
	rangeInt3 := NewRangeInt(0, 1)
	require.False(t, rangeInt1.IsIntersect(rangeInt2))
	require.False(t, rangeInt1.IsIntersect(rangeInt3))
}

func TestToSliceNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(5, 10)
	require.Equal(t, []int{5, 6, 7, 8, 9, 10}, rangeInt1.ToSlice())
}

func TestToSliceEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	require.Equal(t, []int{}, rangeInt1.ToSlice())
}

func TestMinimumNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(5, 10)
	ans, res := rangeInt1.Minimum()
	require.Equal(t, 5, ans)
	require.True(t, res)
}

func TestMinimumEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	ans, res := rangeInt1.Minimum()
	require.Equal(t, 0, ans)
	require.False(t, res)
}

func TestMaximumNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(5, 10)
	ans, res := rangeInt1.Maximum()
	require.Equal(t, 10, ans)
	require.True(t, res)
}

func TestMaximumEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	ans, res := rangeInt1.Maximum()
	require.Equal(t, 0, ans)
	require.False(t, res)
}

func TestStringNormal(t *testing.T) {
	rangeInt1 := NewRangeInt(5, 10)
	require.Equal(t, "[5,10]", rangeInt1.String())
}

func TestStringEmpty(t *testing.T) {
	rangeInt1 := NewRangeInt(1, 0)
	require.Equal(t, "", rangeInt1.String())
}
