package rangeI

import "fmt"

type RangeInt interface {
	Length() int
	Intersect(other RangeInt)
	Union(other RangeInt) bool
	IsEmpty() bool
	ContainsInt(i int) bool
	ContainsRange(other RangeInt) bool
	IsIntersect(other RangeInt) bool
	ToSlice() []int
	Minimum() (int, bool)
	Maximum() (int, bool)
	String() string
}

type Interval struct {
	left  int
	right int
}

func (interval *Interval) Length() int {
	if !interval.IsEmpty() {
		return interval.right - interval.left + 1
	} else {
		return 0
	}
}

func (interval *Interval) Intersect(other RangeInt) {
	left2, _ := other.Minimum()
	right2, _ := other.Maximum()

	if !interval.IsIntersect(other) {
		interval.left = 1
		interval.right = 0
		return
	}

	if interval.left < left2 {
		interval.left = left2
	}

	if interval.right > right2 {
		interval.right = right2
	}
}

func (interval *Interval) Union(other RangeInt) bool {
	if other.IsEmpty() {
		return true
	}

	left2, _ := other.Minimum()
	right2, _ := other.Maximum()

	if interval.IsEmpty() {
		interval.left = left2
		interval.right = right2
		return true
	}

	if (interval.left > right2+1) || (interval.right+1 < left2) {
		return false
	}

	if interval.left > left2 {
		interval.left = left2
	}

	if interval.right < right2 {
		interval.right = right2
	}
	return true
}

func (interval *Interval) IsEmpty() bool {
	return interval.left > interval.right
}

func (interval *Interval) ContainsInt(i int) bool {
	return interval.left <= i && interval.right >= i
}

func (interval *Interval) ContainsRange(other RangeInt) bool {
	if other.IsEmpty() {
		return true
	}

	left2, _ := other.Minimum()
	right2, _ := other.Maximum()

	if interval.left <= left2 && interval.right >= right2 {
		return true
	}
	return false
}

func (interval *Interval) IsIntersect(other RangeInt) bool {
	left2, ok := other.Minimum()
	right2, _ := other.Maximum()

	if interval.IsEmpty() {
		return false
	}
	return !(!ok || (interval.left > right2) || (interval.right < left2))
}

func (interval *Interval) ToSlice() []int {
	a := []int{}
	if interval.IsEmpty() {
		return a
	}
	for i, index := interval.left, 0; i <= interval.right; i, index = i+1, index+1 {
		//a[index] = i
		a = append(a, i)
	}
	return a
}

func (interval *Interval) Minimum() (int, bool) {
	if !interval.IsEmpty() {
		return interval.left, true
	} else {
		return 0, false
	}
}

func (interval *Interval) Maximum() (int, bool) {
	if !interval.IsEmpty() {
		return interval.right, true
	} else {
		return 0, false
	}
}

func (interval *Interval) String() string {
	if interval.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("[%d,%d]", interval.left, interval.right)
}

func NewRangeInt(a, b int) *Interval {
	return &Interval{
		left:  a,
		right: b,
	}
}
